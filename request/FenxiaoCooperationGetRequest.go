package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type FenxiaoCooperationGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewFenxiaoCooperationGetRequest() (req *FenxiaoCooperationGetRequest) {
	request := topsdk.Request{MethodName: "taobao.fenxiao.cooperation.get", Params: make(map[string]interface{}, 6)}
	req = &FenxiaoCooperationGetRequest{
		Request: &request,
	}
	return
}

// 合作结束时间yyyy-MM-dd HH:mm:ss
func (req *FenxiaoCooperationGetRequest) SetEndDate(endDate string) {
	req.Request.Params["end_date"] = endDate
}

func (req *FenxiaoCooperationGetRequest) GetEndDate() string {
	endDate, found := req.Request.Params["end_date"]
	if found {
		return endDate.(string)
	}
	return ""
}

// 合作开始时间yyyy-MM-dd HH:mm:ss
func (req *FenxiaoCooperationGetRequest) SetStartDate(startDate string) {
	req.Request.Params["start_date"] = startDate
}

func (req *FenxiaoCooperationGetRequest) GetStartDate() string {
	startDate, found := req.Request.Params["start_date"]
	if found {
		return startDate.(string)
	}
	return ""
}

// 合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止)
func (req *FenxiaoCooperationGetRequest) SetStatus(status string) {
	req.Request.Params["status"] = status
}

func (req *FenxiaoCooperationGetRequest) GetStatus() string {
	status, found := req.Request.Params["status"]
	if found {
		return status.(string)
	}
	return ""
}

// 分销方式：AGENT(代销) 、DEALER（经销）
func (req *FenxiaoCooperationGetRequest) SetTradeType(tradeType string) {
	req.Request.Params["trade_type"] = tradeType
}

func (req *FenxiaoCooperationGetRequest) GetTradeType() string {
	tradeType, found := req.Request.Params["trade_type"]
	if found {
		return tradeType.(string)
	}
	return ""
}

//页码（大于0的整数，默认1）
func (req *FenxiaoCooperationGetRequest) SetPageNo(page_no uint16) {
	req.Request.Params["page_no"] = page_no
}

func (req *FenxiaoCooperationGetRequest) GetPageNo() uint16 {
	page_no, found := req.Request.Params["page_no"]
	if found {
		return page_no.(uint16)
	}
	return 0
}

// 每页记录数（默认20，最大50）
func (req *FenxiaoCooperationGetRequest) SetPageSize(page_size uint16) {
	req.Request.Params["page_size"] = page_size
}

func (req *FenxiaoCooperationGetRequest) GetPageSize() uint16 {
	page_size, found := req.Request.Params["page_size"]
	if found {
		return page_size.(uint16)
	}
	return 0
}
