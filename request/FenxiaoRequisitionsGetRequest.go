package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type FenxiaoRequisitionsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewFenxiaoRequisitionsGetRequest() (req *FenxiaoRequisitionsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.fenxiao.requisitions.get", Params: make(map[string]interface{}, 5)}
	req = &FenxiaoRequisitionsGetRequest{
		Request: &request,
	}
	return
}

// 申请结束时间yyyy-MM-dd
func (req *FenxiaoRequisitionsGetRequest) SetApplyEnd(applyEnd string) {
	req.Request.Params["apply_end"] = applyEnd
}

func (req *FenxiaoRequisitionsGetRequest) GetApplyEnd() string {
	applyEnd, found := req.Request.Params["apply_end"]
	if found {
		return applyEnd.(string)
	}
	return ""
}

// 申请开始时间yyyy-MM-dd
func (req *FenxiaoRequisitionsGetRequest) SetApplyStart(applyStart string) {
	req.Request.Params["apply_start"] = applyStart
}

func (req *FenxiaoRequisitionsGetRequest) GetApplyStart() string {
	applyStart, found := req.Request.Params["apply_start"]
	if found {
		return applyStart.(string)
	}
	return ""
}

//页码（大于0的整数，默认1）
func (req *FenxiaoRequisitionsGetRequest) SetPageNo(page_no uint16) {
	req.Request.Params["page_no"] = page_no
}

func (req *FenxiaoRequisitionsGetRequest) GetPageNo() uint16 {
	page_no, found := req.Request.Params["page_no"]
	if found {
		return page_no.(uint16)
	}
	return 0
}

// 每页记录数（默认20，最大50）
func (req *FenxiaoRequisitionsGetRequest) SetPageSize(page_size uint16) {
	req.Request.Params["page_size"] = page_size
}

func (req *FenxiaoRequisitionsGetRequest) GetPageSize() uint16 {
	page_size, found := req.Request.Params["page_size"]
	if found {
		return page_size.(uint16)
	}
	return 0
}

// 申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期）
func (req *FenxiaoRequisitionsGetRequest) SetStatus(status uint8) {
	req.Request.Params["status"] = status
}

func (req *FenxiaoRequisitionsGetRequest) GetStatus() uint8 {
	status, found := req.Request.Params["status"]
	if found {
		return status.(uint8)
	}
	return 0
}
