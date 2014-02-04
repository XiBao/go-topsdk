package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaRptCustbaseGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaRptCustbaseGetRequest() (req *SimbaRptCustbaseGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.rpt.custbase.get", Params:make(map[string]interface{}, 8)}
    req = &SimbaRptCustbaseGetRequest {
        Request: &request,
    }
    return
}

//主人昵称
func (req *SimbaRptCustbaseGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaRptCustbaseGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}

//开始时间
func (req *SimbaRptCustbaseGetRequest) SetStartTime(start_time string) {
    req.Request.Params["start_time"] = start_time
}

func (req *SimbaRptCustbaseGetRequest) GetStartTime() string {
    start_time, found := req.Request.Params["start_time"]
    if found {
        return start_time.(string)
    }
    return ""
}

//结束时间
func (req *SimbaRptCustbaseGetRequest) SetEndTime(end_time string) {
    req.Request.Params["end_time"] = end_time
}

func (req *SimbaRptCustbaseGetRequest) GetEndTime() string {
    end_time, found := req.Request.Params["end_time"]
    if found {
        return end_time.(string)
    }
    return ""
}

/** 
 * 报表类型（搜索：SEARCH,类目出价：CAT,
定向投放：NOSEARCH）可以一次取多个例如：SEARCH,CAT
 **/
func (req *SimbaRptCustbaseGetRequest) SetSearchType(search_type string) {
    req.Request.Params["search_type"] = search_type
}

func (req *SimbaRptCustbaseGetRequest) GetSearchType() string {
    search_type, found := req.Request.Params["search_type"]
    if found {
        return search_type.(string)
    }
    return ""
}

//数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2
func (req *SimbaRptCustbaseGetRequest) SetSource(source string) {
    req.Request.Params["source"] = source
}

func (req *SimbaRptCustbaseGetRequest) GetSource() string {
    source, found := req.Request.Params["source"]
    if found {
        return source.(string)
    }
    return ""
}

// 权限校验参数
func (req *SimbaRptCustbaseGetRequest) SetSubwayToken(subway_token string) {
    req.Request.Params["subway_token"] = subway_token
}

func (req *SimbaRptCustbaseGetRequest) GetSubwayToken() string {
    subway_token, found := req.Request.Params["subway_token"]
    if found {
        return subway_token.(string)
    }
    return ""
}

//页码，从1开始
func (req *SimbaRptCustbaseGetRequest) SetPageNo(page_no uint16) {
    req.Request.Params["page_no"] = page_no
}

func (req *SimbaRptCustbaseGetRequest) GetPageNo() uint16 {
    page_no, found := req.Request.Params["page_no"]
    if found {
        return page_no.(uint16)
    }
    return 0
}

// 每页大小
func (req *SimbaRptCustbaseGetRequest) SetPageSize(page_size uint16) {
    req.Request.Params["page_size"] = page_size
}

func (req *SimbaRptCustbaseGetRequest) GetPageSize() uint16 {
    page_size, found := req.Request.Params["page_size"]
    if found {
        return page_size.(uint16)
    }
    return 0
}