package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type TraderatesGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewTraderatesGetRequest() (req *TraderatesGetRequest) {
    request := topsdk.Request{MethodName:"taobao.traderates.get", Params:make(map[string]interface{}, 10)}
    req = &TraderatesGetRequest {
        Request: &request,
    }
    return
}

//需返回的字段列表。可选值：TradeRate 结构中的所有字段，多个字段之间用“,”分隔
func (req *TraderatesGetRequest) SetFields(fields string) {
    req.Request.Params["fields"] = fields
}

func (req *TraderatesGetRequest) GetFields() string {
    fields, found := req.Request.Params["fields"]
    if found {
        return fields.(string)
    }
    return ""
}

// 评价结束时间。如果只输入结束时间，那么全部返回所有评价数据。
func (req *TraderatesGetRequest) SetEndDate(endDate string) {
    req.Request.Params["end_date"] = endDate
}

func (req *TraderatesGetRequest) GetEndDate() string {
    endDate, found := req.Request.Params["end_date"]
    if found {
        return endDate.(string)
    }
    return ""
}

// 评价开始时。如果只输入开始时间，那么能返回开始时间之后的评价数据。
func (req *TraderatesGetRequest) SetStartDate(startDate string) {
    req.Request.Params["start_date"] = startDate
}

func (req *TraderatesGetRequest) GetStartDate() string {
    startDate, found := req.Request.Params["start_date"]
    if found {
        return startDate.(string)
    }
    return ""
}

func (req *TraderatesGetRequest) SetTid(tid uint64) {
    req.Request.Params["tid"] = tid
}

func (req *TraderatesGetRequest) GetTid() uint64 {
    tid, found := req.Request.Params["tid"]
    if found {
        return tid.(uint64)
    }
    return 0
}

func (req *TraderatesGetRequest) SetNumIid(numiid uint64) {
    req.Request.Params["num_iid"] = numiid
}

func (req *TraderatesGetRequest) GetNumIid() uint64 {
    numiid, found := req.Request.Params["num_iid"]
    if found {
        return numiid.(uint64)
    }
    return 0
}

func (req *TraderatesGetRequest) SetPageNo(pageNo uint64) {
    req.Request.Params["page_no"] = pageNo
}

func (req *TraderatesGetRequest) GetPageNo() uint64 {
    pageNo, found := req.Request.Params["page_no"]
    if found {
        return pageNo.(uint64)
    }
    return 0
}

func (req *TraderatesGetRequest) SetPageSize(pageSize uint64) {
    req.Request.Params["page_size"] = pageSize
}

func (req *TraderatesGetRequest) GetPageSize() uint64 {
    pageSize, found := req.Request.Params["page_size"]
    if found {
        return pageSize.(uint64)
    }
    return 0
}

func (req *TraderatesGetRequest) SetRateType(rateType string) {
    req.Request.Params["rate_type"] = rateType
}

func (req *TraderatesGetRequest) GetRateType() string {
    rateType, found := req.Request.Params["rate_type"]
    if found {
        return rateType.(string)
    }
    return ""
}

func (req *TraderatesGetRequest) SetResult(result string) {
    req.Request.Params["result"] = result 
}

func (req *TraderatesGetRequest) GetResult() string {
    result, found := req.Request.Params["result"]
    if found {
        return result.(string)
    }
    return ""
}

func (req *TraderatesGetRequest) SetRole(role string) {
    req.Request.Params["role"] = role
}

func (req *TraderatesGetRequest) GetRole() string {
    role, found := req.Request.Params["role"]
    if found {
        return role.(string)
    }
    return ""
}

func (req *TraderatesGetRequest) SetUseHasNext(useHasNext bool) {
    if useHasNext {
        req.Request.Params["use_has_next"] = "true"
    }else{
        req.Request.Params["use_has_next"] = "false"
    }
}

func (req *TraderatesGetRequest) GetUseHasNext() bool {
    useHasNext, found := req.Request.Params["use_has_next"]
    if found {
        return useHasNext.(string) == "true"
    }
    return false
}


