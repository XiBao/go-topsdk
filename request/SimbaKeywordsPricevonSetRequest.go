package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaKeywordsPricevonSetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaKeywordsPricevonSetRequest() (req *SimbaKeywordsPricevonSetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.keywords.pricevon.set", Params:make(map[string]interface{}, 3)}
    req = &SimbaKeywordsPricevonSetRequest {
        Request: &request,
    }
    return
}

/** 
 * 关键词ID，出价和匹配方式json字符串，keywordId:词ID，整数。maxPrice：价格，是整数，以“分”为单位，不能小于5，不能大于日限额,当使用默认出价时必须将这个值设置为0。; isDefaultPrice：是否使用默认出价，只能是0，1(0代表不使用，1代表使用)。matchscope只能是1,2,4（1代表精确匹配，2代表子串匹配，4代表广泛匹配）
 **/
func (req *SimbaKeywordsPricevonSetRequest) SetKeywordidPrices(keywordid_prices string) {
    req.Request.Params["keywordid_prices"] = keywordid_prices
}

func (req *SimbaKeywordsPricevonSetRequest) GetKeywordidPrices() string {
    keywordid_prices, found := req.Request.Params["keywordid_prices"]
    if found {
        return keywordid_prices.(string)
    }
    return ""
}

//主人昵称
func (req *SimbaKeywordsPricevonSetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaKeywordsPricevonSetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
