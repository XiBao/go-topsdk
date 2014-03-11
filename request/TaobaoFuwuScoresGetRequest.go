package TopRequest

// 根据日期、查询appkey对应服务评价，每次调用只能查询某一天服务评价信息，可设置分页查询，页大小最大为100，非实时接口，延迟时间为30分钟

import (
    "github.com/XiBao/topsdk"
)

type TaobaoFuwuScoresGetRequest struct {
    Request *topsdk.Request
}

func NewTaobaoFuwuScoresGetRequest() (req *TaobaoFuwuScoresGetRequest) {
    request := topsdk.Request{MethodName:"taobao.fuwu.scores.get", Params:make(map[string]interface{}, 3)}
    req = &TaobaoFuwuScoresGetRequest {
        Request: &request,
    }
    return
}

func (req *TaobaoFuwuScoresGetRequest) SetCurrentPage(currentPage uint16) {
    req.Request.Params["current_page"] = currentPage
}

func (req *TaobaoFuwuScoresGetRequest) GetCurrentPage() uint16 {
    currentPage, found := req.Request.Params["current_page"]
    if found {
        return currentPage.(uint16)
    }
    return 0
}

func (req *TaobaoFuwuScoresGetRequest) SetPageSize(pageSize uint16) {
    req.Request.Params["page_size"] = pageSize
}

func (req *TaobaoFuwuScoresGetRequest) GetPageSize() uint16 {
    pageSize, found := req.Request.Params["page_size"]
    if found {
        return pageSize.(uint16)
    }
    return 0
}

func (req *TaobaoFuwuScoresGetRequest) SetDate(date string) {
    req.Request.Params["date"] = date
}

func (req *TaobaoFuwuScoresGetRequest) GetDate() string {
    date, found := req.Request.Params["date"]
    if found {
        return date.(string)
    }
    return ""
}
