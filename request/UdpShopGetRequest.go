package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type UdpShopGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewUdpShopGetRequest() (req *UdpShopGetRequest) {
    request := topsdk.Request{MethodName:"taobao.udp.shop.get", Params:make(map[string]interface{}, 5)}
    req = &UdpShopGetRequest {
        Request: &request,
    }
    return
}

// 地区ID
func (req *UdpShopGetRequest) SetArea(area uint64) {
    req.Request.Params["area"] = area
}

func (req *UdpShopGetRequest) GetArea() uint64 {
    area, found := req.Request.Params["area"]
    if found {
        return area.(uint64)
    }
    return 0
}

//开始时间
func (req *UdpShopGetRequest) SetBeginTime(beginTime string) {
    req.Request.Params["begin_time"] = beginTime 
}

func (req *UdpShopGetRequest) GetBeginTime() string {
    beginTime, found := req.Request.Params["begin_time"]
    if found {
        return beginTime.(string)
    }
    return ""
}

//结束时间
func (req *UdpShopGetRequest) SetEndTime(endTime string) {
    req.Request.Params["end_time"] = endTime 
}

func (req *UdpShopGetRequest) GetEndTime() string {
    endTime, found := req.Request.Params["end_time"]
    if found {
        return endTime.(string)
    }
    return ""
}

//指标ID
func (req *UdpShopGetRequest) SetFields(fields string) {
    req.Request.Params["fields"] = fields 
}

func (req *UdpShopGetRequest) GetFields() string {
    fields, found := req.Request.Params["fields"]
    if found {
        return fields.(string)
    }
    return ""
}

//保留字段
func (req *UdpShopGetRequest) SetParameters(parameters string) {
    req.Request.Params["parameters"] = parameters 
}

func (req *UdpShopGetRequest) GetParameters() string {
    parameters, found := req.Request.Params["parameters"]
    if found {
        return parameters.(string)
    }
    return ""
}
