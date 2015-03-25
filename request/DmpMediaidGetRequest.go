package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpMediaidGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpMediaidGetRequest() (req *DmpMediaidGetRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.mediaid.get", Params: make(map[string]interface{}, 3)}
	req = &DmpMediaidGetRequest{
		Request: &request,
	}
	return
}

// 可传入多个ID,逗号分隔
func (req *DmpMediaidGetRequest) SetIds(ids string) {
	req.Request.Params["ids"] = ids
}

func (req *DmpMediaidGetRequest) GetIds() string {
	ids, found := req.Request.Params["ids"]
	if found {
		return ids.(string)
	}
	return ""
}

// 支持idfa, mac, imei
func (req *DmpMediaidGetRequest) SetType(atype string) {
	req.Request.Params["type"] = atype
}

func (req *DmpMediaidGetRequest) GetType() string {
	atype, found := req.Request.Params["type"]
	if found {
		return atype.(string)
	}
	return ""
}

// 是否加密，目前仅支持不加密数据传入
func (req *DmpMediaidGetRequest) SetEncrypt(encrypt uint8) {
	req.Request.Params["encrypt"] = encrypt
}

func (req *DmpMediaidGetRequest) GetEncrypt() uint8 {
	encrypt, found := req.Request.Params["encrypt"]
	if found {
		return encrypt.(uint8)
	}
	return 0
}
