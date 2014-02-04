package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type UdpItemGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewUdpItemGetRequest() (req *UdpItemGetRequest) {
    request := topsdk.Request{MethodName:"taobao.udp.item.get", Params:make(map[string]interface{}, 7)}
    req = &UdpItemGetRequest {
        Request: &request,
    }
    return
}

// 地区ID
func (req *UdpItemGetRequest) SetArea(area uint64) {
    req.Request.Params["area"] = area
}

func (req *UdpItemGetRequest) GetArea() uint64 {
    area, found := req.Request.Params["area"]
    if found {
        return area.(uint64)
    }
    return 0
}

//开始时间
func (req *UdpItemGetRequest) SetBeginTime(beginTime string) {
    req.Request.Params["begin_time"] = beginTime 
}

func (req *UdpItemGetRequest) GetBeginTime() string {
    beginTime, found := req.Request.Params["begin_time"]
    if found {
        return beginTime.(string)
    }
    return ""
}

//结束时间
func (req *UdpItemGetRequest) SetEndTime(endTime string) {
    req.Request.Params["end_time"] = endTime 
}

func (req *UdpItemGetRequest) GetEndTime() string {
    endTime, found := req.Request.Params["end_time"]
    if found {
        return endTime.(string)
    }
    return ""
}

//指标ID
func (req *UdpItemGetRequest) SetFields(fields string) {
    req.Request.Params["fields"] = fields 
}

func (req *UdpItemGetRequest) GetFields() string {
    fields, found := req.Request.Params["fields"]
    if found {
        return fields.(string)
    }
    return ""
}

//商品ID
func (req *UdpItemGetRequest) SetItemId(itemId uint64) {
    req.Request.Params["itemid"] = itemId 
}

func (req *UdpItemGetRequest) GetItemId() uint64 {
    itemId, found := req.Request.Params["itemid"]
    if found {
        return itemId.(uint64)
    }
    return 0
}

//来源编号
func (req *UdpItemGetRequest) SetSource(source uint) {
    req.Request.Params["source"] = source 
}

func (req *UdpItemGetRequest) GetSource() uint {
    source, found := req.Request.Params["source"]
    if found {
        return source.(uint)
    }
    return 0
}

//保留字段
func (req *UdpItemGetRequest) SetParameters(parameters string) {
    req.Request.Params["parameters"] = parameters 
}

func (req *UdpItemGetRequest) GetParameters() string {
    parameters, found := req.Request.Params["parameters"]
    if found {
        return parameters.(string)
    }
    return ""
}
