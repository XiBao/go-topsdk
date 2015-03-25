package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpDpeffectfeedbackGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpDpeffectfeedbackGetRequest() (req *DmpDpeffectfeedbackGetRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.dpeffectfeedback.get", Params: make(map[string]interface{}, 5)}
	req = &DmpDpeffectfeedbackGetRequest{
		Request: &request,
	}
	return
}

// 回流效果数据的日期
func (req *DmpDpeffectfeedbackGetRequest) SetThedate(date string) {
	req.Request.Params["thedate"] = date
}

func (req *DmpDpeffectfeedbackGetRequest) GetThedate() string {
	date, found := req.Request.Params["thedate"]
	if found {
		return date.(string)
	}
	return ""
}

// 用户淘宝ID
func (req *DmpDpeffectfeedbackGetRequest) SetTbUserId(id uint64) {
	req.Request.Params["tb_user_id"] = id
}

func (req *DmpDpeffectfeedbackGetRequest) GetTbUserId() uint64 {
	id, found := req.Request.Params["tb_user_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 需要请求的效果字段，英文逗号隔开
func (req *DmpDpeffectfeedbackGetRequest) SetColumns(columns string) {
	req.Request.Params["columns"] = columns
}

func (req *DmpDpeffectfeedbackGetRequest) GetColumns() string {
	columns, found := req.Request.Params["columns"]
	if found {
		return columns.(string)
	}
	return ""
}

// 返回数据的第N页
func (req *DmpDpeffectfeedbackGetRequest) SetPage(page uint) {
	req.Request.Params["page"] = page
}

func (req *DmpDpeffectfeedbackGetRequest) GetPage() uint {
	page, found := req.Request.Params["page"]
	if found {
		return page.(uint)
	}
	return 0
}

// 返回数据，每页M行
func (req *DmpDpeffectfeedbackGetRequest) SetPageSize(pageSize uint) {
	req.Request.Params["page_size"] = pageSize
}

func (req *DmpDpeffectfeedbackGetRequest) GetPageSize() uint {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(uint)
	}
	return 0
}
