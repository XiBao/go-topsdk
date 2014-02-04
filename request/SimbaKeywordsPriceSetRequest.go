package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaKeywordsPriceSetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaKeywordsPriceSetRequest() (req *SimbaKeywordsPriceSetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.keywords.price.set", Params:make(map[string]interface{}, 3)}
    req = &SimbaKeywordsPriceSetRequest {
        Request: &request,
    }
    return
}

// 推广组id
func (req *SimbaKeywordsPriceSetRequest) SetAdgroupId(adgroup_id uint64) {
    req.Request.Params["adgroup_id"] = adgroup_id
}

func (req *SimbaKeywordsPriceSetRequest) GetAdgroupId() uint64 {
    adgroup_id, found := req.Request.Params["adgroup_id"]
    if found {
        return adgroup_id.(uint64)
    }
    return 0
}

/** 
 * 关键词Id出价字符串和匹配方式字符串数组，最多200个;
每个字符串：keywordId+  ”^^”+price+”^^”+matchscope；
Price是整数，以“分”为单位，不能小于5，不能大于日限额; 如果该词为无展现词，出价需要大于原来出价，才会生效。
price为0则设置为使用默认出价；
matchscope只能是1,2,4 (1代表精确匹配，2代表子串匹配，4代表广泛匹配) 可不传
例如102232^^85，102231^^82^^4
 **/
func (req *SimbaKeywordsPriceSetRequest) SetKeywordidPrices(keywordid_prices string) {
    req.Request.Params["keywordid_prices"] = keywordid_prices
}

func (req *SimbaKeywordsPriceSetRequest) GetKeywordidPrices() string {
    keywordid_prices, found := req.Request.Params["keywordid_prices"]
    if found {
        return keywordid_prices.(string)
    }
    return ""
}

//主人昵称
func (req *SimbaKeywordsPriceSetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaKeywordsPriceSetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
