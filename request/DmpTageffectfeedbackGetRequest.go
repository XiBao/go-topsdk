package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpTageffectfeedbackGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpTageffectfeedbackGetRequest() (req *DmpTageffectfeedbackGetRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.tageffectfeedback.get", Params: make(map[string]interface{}, 5)}
	req = &DmpTageffectfeedbackGetRequest{
		Request: &request,
	}
	return
}

// 回流效果数据的日期
func (req *DmpTageffectfeedbackGetRequest) SetThedate(date string) {
	req.Request.Params["thedate"] = date
}

func (req *DmpTageffectfeedbackGetRequest) GetThedate() string {
	date, found := req.Request.Params["thedate"]
	if found {
		return date.(string)
	}
	return ""
}

// 用户淘宝ID
func (req *DmpTageffectfeedbackGetRequest) SetTbUserId(id uint64) {
	req.Request.Params["tb_user_id"] = id
}

func (req *DmpTageffectfeedbackGetRequest) GetTbUserId() uint64 {
	id, found := req.Request.Params["tb_user_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 需要请求的效果字段，英文逗号隔开
func (req *DmpTageffectfeedbackGetRequest) SetColumns(columns string) {
	req.Request.Params["columns"] = columns
}

func (req *DmpTageffectfeedbackGetRequest) GetColumns() string {
	columns, found := req.Request.Params["columns"]
	if found {
		return columns.(string)
	}
	return ""
}

// 返回数据的第N页
func (req *DmpTageffectfeedbackGetRequest) SetPage(page uint) {
	req.Request.Params["page"] = page
}

func (req *DmpTageffectfeedbackGetRequest) GetPage() uint {
	page, found := req.Request.Params["page"]
	if found {
		return page.(uint)
	}
	return 0
}

// 返回数据，每页M行
func (req *DmpTageffectfeedbackGetRequest) SetPageSize(pageSize uint) {
	req.Request.Params["page_size"] = pageSize
}

func (req *DmpTageffectfeedbackGetRequest) GetPageSize() uint {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(uint)
	}
	return 0
}
