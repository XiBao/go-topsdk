package topsdk

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/bububa/ljson"
	"github.com/bububa/x2j"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	MethodName string
	Format     string
	Params     map[string]interface{}
}

type Client struct {
	AppKey    string
	SecretKey string
}

//create new client
func NewClient(appKey string, secretKey string) (c *Client) {
	c = &Client{
		AppKey:    appKey,
		SecretKey: secretKey,
	}
	return
}

func (c *Client) Execute(req *Request, session ...string) (result interface{}, err error) {
	sysParams := make(map[string]string, len(req.Params)+9)
	sysParams["app_key"] = c.AppKey
	sysParams["v"] = API_VERSION
	switch req.Format {
	case XML_FORMAT:
		sysParams["format"] = XML_FORMAT
	default:
		sysParams["format"] = JSON_FORMAT
	}
	sysParams["sign_method"] = SIGN_METHOD
	sysParams["method"] = req.MethodName
	sysParams["timestamp"] = time.Now().Local().Format("2006-01-02 15:04:05")
	sysParams["partner_id"] = SDK_VERSION
	if len(session) > 0 && session[0] != "" {
		sysParams["session"] = session[0]
	}
	for k, v := range req.Params {
		switch v.(type) {
		case string:
			sysParams[k] = v.(string)
		case uint64:
			sysParams[k] = fmt.Sprintf("%d", v.(uint64))
		case uint32:
			sysParams[k] = fmt.Sprintf("%d", v.(uint32))
		case uint16:
			sysParams[k] = fmt.Sprintf("%d", v.(uint16))
		case uint8:
			sysParams[k] = fmt.Sprintf("%d", v.(uint8))
		case uint:
			sysParams[k] = fmt.Sprintf("%d", v.(uint))
		case float32:
			sysParams[k] = fmt.Sprintf("%f", v.(float32))
		case float64:
			sysParams[k] = fmt.Sprintf("%f", v.(float64))
		case bool:
			if v.(bool) {
				sysParams[k] = "true"
			} else {
				sysParams[k] = "false"
			}
		}
	}
	//log.Println(sysParams)
	rawSign := c.GenerateRawSign(sysParams)
	sysParams["sign"] = c.GenerateSign(rawSign)
	values := url.Values{}
	for k, v := range sysParams {
		values.Add(k, v)
	}
	//reqUrl := GATEWAY_URL + "?" + values.Encode()
	response, e := http.PostForm(GATEWAY_URL, values)
	if e != nil {
		return nil, Error{Code: 0, Msg: "HTTP Response Error"}
	}

	defer response.Body.Close()
	body, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return nil, Error{Code: 0, Msg: fmt.Sprintf("ReadAll on response.Body: %v", e)}
	}

	j := make(map[string]interface{})
	switch sysParams["format"] {
	case JSON_FORMAT:
		/*regex, _ := regexp.Compile("\",,\"")
		res := regex.ReplaceAllString(string(body), "\",\"")
		regex, _ = regexp.Compile("\"tid\":(\\d+)")
		res = regex.ReplaceAllString(res, "\"tid\":\"$1\"")
		regex, _ = regexp.Compile("\"oid\":(\\d+)")
		res = regex.ReplaceAllString(res, "\"oid\":\"$1\"")
		regex, _ = regexp.Compile("\"num_iid\":(\\d+)")
		res = regex.ReplaceAllString(res, "\"num_iid\":\"$1\"")
		regex, _ = regexp.Compile("\"order_id\":(\\d+)")
		res = regex.ReplaceAllString(res, "\"order_id\":\"$1\"")
		regex, _ = regexp.Compile("\"biz_order_id\":(\\d+)")
		res = regex.ReplaceAllString(res, "\"biz_order_id\":\"$1\"")
		regex, _ = regexp.Compile("\"keyword_id\":(\\d+)")
		res = regex.ReplaceAllString(res, "\"keyword_id\":\"$1\"")
		regex, _ = regexp.Compile("\"adgroup_id\":(\\d+)")
		res = regex.ReplaceAllString(res, "\"adgroup_id\":\"$1\"")
		regex, _ = regexp.Compile("\"campaign_id\":(\\d+)")
		res = regex.ReplaceAllString(res, "\"campaign_id\":\"$1\"")*/
		res := string(body)
		res = strings.Replace(res, "\n", "", -1)
		res = strings.Replace(res, "\r", "", -1)
		res = strings.Replace(res, "\t", "", -1)
		e = ljson.Unmarshal([]byte(res), &j)
	case XML_FORMAT:
		j, e = x2j.DocToMap(string(body))
	}
	if e != nil {
		if redirectUrl, e2 := response.Location(); e2 != nil {
			if strings.Contains(string(body), "403 Forbidden") {
				log.Printf("403 Forbidden\n", values)
				return string(body), Error{Code: 0, Msg: "403 Forbidden", SubMsg: "403 Forbidden"}
			}
			log.Printf("HTTP_RESPONSE: (%s) %s\n", GATEWAY_URL+"?"+values.Encode(), string(body))
			return string(body), Error{Code: 0, Msg: "Decode error", SubMsg: e.Error()}
		} else {
			log.Printf("302 Redirect: %s, REQUEST: %s, %v, %s\n", redirectUrl.String(), GATEWAY_URL, values, rawSign)
			return string(body), Error{Code: 0, Msg: "302 Redirect", SubMsg: redirectUrl.String()}
		}
	}
	resp := make(map[string]interface{})
	for _, v := range j {
		resp = v.(map[string]interface{})
	}
	//log.Println(resp)
	err = c.CheckError(resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *Client) CheckError(v map[string]interface{}) (err error) {
	_, found := v["code"]
	if found {
		err = Error{Code: Uint8Assert(v["code"]), Msg: v["msg"].(string)}
		subCode, found := v["sub_code"]
		if found {
			err = Error{Code: Uint8Assert(v["code"]), SubCode: subCode.(string), Msg: v["msg"].(string)}
			subMsg, found := v["sub_msg"]
			if found {
				err = Error{Code: Uint8Assert(v["code"]), SubCode: subCode.(string), Msg: v["msg"].(string), SubMsg: subMsg.(string)}
			}
		}
	}
	return
}

func (c *Client) GenerateRawSign(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		if 64 == key[0] {
			continue
		}
		keys = append(keys, key)
	}
	sort.Strings(keys)
	stringToBeSigned := c.SecretKey
	for _, k := range keys {
		stringToBeSigned += k + params[k]
	}
	stringToBeSigned += c.SecretKey
	return stringToBeSigned
}

func (c *Client) GenerateSign(stringToBeSigned string) string {
	h := md5.New()
	io.WriteString(h, stringToBeSigned)
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func KeywordUstr(str string) string {
	c1 := []string{" ", "男式", "男士", "男款", "女式", "女士", "女款", "夏季", "夏天", "冬季", "冬天", "春季", "春天", "秋季", "秋天"}
	c2 := []string{"", "男", "男", "男", "女", "女", "女", "夏", "夏", "冬", "冬", "春", "春", "秋", "秋"}
	for i, c := range c1 {
		str = strings.Replace(str, c, c2[i], -1)
	}
	chars := strings.Split(strings.ToLower(str), "")
	sort.Strings(chars)
	orderedStr := strings.Join(chars, "")
	h := md5.New()
	io.WriteString(h, orderedStr)
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}

func Sha1(str string) string {
	h := sha1.New()
	io.WriteString(h, str)
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}

func Uint8Assert(val interface{}) uint8 {
	switch val.(type) {
	case uint:
		return uint8(val.(uint))
	case int:
		return uint8(val.(int))
	case int64:
		return uint8(val.(int64))
	case int32:
		return uint8(val.(int32))
	case int16:
		return uint8(val.(int16))
	case int8:
		return uint8(val.(int8))
	case uint8:
		return val.(uint8)
	case uint16:
		return uint8(val.(uint16))
	case uint32:
		return uint8(val.(uint32))
	case uint64:
		return uint8(val.(uint64))
	case float64:
		return uint8(val.(float64))
	case float32:
		return uint8(val.(float32))
	case nil:
		return uint8(0)
	case string:
		n, e := strconv.ParseUint(val.(string), 10, 64)
		if e != nil {
			return uint8(0)
		}
		return uint8(n)
	case bool:
		switch val.(bool) {
		case true:
			return uint8(1)
		case false:
			return uint8(0)
		}
	}
	return val.(uint8)
}
