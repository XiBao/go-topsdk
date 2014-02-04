package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaKeywordsRecommendGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaKeywordsRecommendGetRequest() (req *SimbaKeywordsRecommendGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.keywords.recommend.get", Params:make(map[string]interface{}, 6)}
    req = &SimbaKeywordsRecommendGetRequest {
        Request: &request,
    }
    return
}

// 推广组ID
func (req *SimbaKeywordsRecommendGetRequest) SetAdgroupId(adgroup_id uint64) {
    req.Request.Params["adgroup_id"] = adgroup_id
}

func (req *SimbaKeywordsRecommendGetRequest) GetAdgroupId() uint64 {
    adgroup_id, found := req.Request.Params["adgroup_id"]
    if found {
        return adgroup_id.(uint64)
    }
    return 0
}

//主人昵称
func (req *SimbaKeywordsRecommendGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaKeywordsRecommendGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}

/*
* 排序方式: 
搜索量 search_volume 
市场平均价格 average_price 
相关度 relevance 
不排序 non 
默认为 non
*/
func (req *SimbaKeywordsRecommendGetRequest) SetOrderBy(order_by string) {
    req.Request.Params["order_by"] = order_by
}

func (req *SimbaKeywordsRecommendGetRequest) GetOrderBy() string {
    order_by, found := req.Request.Params["order_by"]
    if found {
        return order_by.(string)
    }
    return ""
}

//返回的第几页数据，默认为1
func (req *SimbaKeywordsRecommendGetRequest) SetPageNo(page_no uint) {
    req.Request.Params["page_no"] = page_no
}

func (req *SimbaKeywordsRecommendGetRequest) GetPageNo() uint {
    page_no, found := req.Request.Params["page_no"]
    if found {
        return page_no.(uint)
    }
    return 0
}

// 返回的每页数据量大小,最大200
func (req *SimbaKeywordsRecommendGetRequest) SetPageSize(page_size uint) {
    req.Request.Params["page_size"] = page_size
}

func (req *SimbaKeywordsRecommendGetRequest) GetItemId() uint {
    page_size, found := req.Request.Params["page_size"]
    if found {
        return page_size.(uint)
    }
    return 0
}

//相关度
func (req *SimbaKeywordsRecommendGetRequest) SetPertinence(pertinence uint) {
    req.Request.Params["pertinence"] = pertinence
}

func (req *SimbaKeywordsRecommendGetRequest) GetPertinence() uint {
    pertinence, found := req.Request.Params["pertinence"]
    if found {
        return pertinence.(uint)
    }
    return 0
}

//搜索量,设置此值后返回的就是大于此搜索量的词列表
func (req *SimbaKeywordsRecommendGetRequest) SetSearch(search uint64) {
    req.Request.Params["search"] = search
}

func (req *SimbaKeywordsRecommendGetRequest) GetSearch() uint64 {
    search, found := req.Request.Params["search"]
    if found {
        return search.(uint64)
    }
    return 0
}