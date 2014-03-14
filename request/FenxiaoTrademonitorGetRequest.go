package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type FenxiaoTrademonitorGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewFenxiaoTrademonitorGetRequest() (req *FenxiaoTrademonitorGetRequest) {
	request := topsdk.Request{MethodName: "taobao.fenxiao.trademonitor.get", Params: make(map[string]interface{}, 7)}
	req = &FenxiaoTrademonitorGetRequest{
		Request: &request,
	}
	return
}

// 结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。
func (req *FenxiaoTrademonitorGetRequest) SetEndCreated(endCreated string) {
	req.Request.Params["end_created"] = endCreated
}

func (req *FenxiaoTrademonitorGetRequest) GetEndCreated() string {
	endCreated, found := req.Request.Params["end_created"]
	if found {
		return endCreated.(string)
	}
	return ""
}

// 起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。
func (req *FenxiaoTrademonitorGetRequest) SetStartCreated(startCreated string) {
	req.Request.Params["start_created"] = startCreated
}

func (req *FenxiaoTrademonitorGetRequest) GetStartCreated() string {
	startCreated, found := req.Request.Params["start_created"]
	if found {
		return startCreated.(string)
	}
	return ""
}

// 多个字段用","分隔。 fields 如果为空：返回所有采购单对象(purchase_orders)字段。 如果不为空：返回指定采购单对象(purchase_orders)字段。例如：trade_monitors.item_title表示只返回item_title
func (req *FenxiaoTrademonitorGetRequest) SetFields(fields string) {
	req.Request.Params["fields"] = fields
}

func (req *FenxiaoTrademonitorGetRequest) GetFields() string {
	fields, found := req.Request.Params["fields"]
	if found {
		return fields.(string)
	}
	return ""
}

// 经销商的淘宝账号
func (req *FenxiaoTrademonitorGetRequest) SetDistributorNick(distributorNick string) {
	req.Request.Params["distributor_nick"] = distributorNick
}

func (req *FenxiaoTrademonitorGetRequest) GetDistributorNick() string {
	distributorNick, found := req.Request.Params["distributor_nick"]
	if found {
		return distributorNick.(string)
	}
	return ""
}

// 产品线ID
func (req *FenxiaoTrademonitorGetRequest) SetProductId(productId uint64) {
	req.Request.Params["product_id"] = productId
}

func (req *FenxiaoTrademonitorGetRequest) GetProductId() uint64 {
	productId, found := req.Request.Params["product_id"]
	if found {
		return productId.(uint64)
	}
	return 0
}

// 页码。（大于0的整数。小于1按1计）
func (req *FenxiaoTrademonitorGetRequest) SetPageNo(page_no uint16) {
	req.Request.Params["page_no"] = page_no
}

func (req *FenxiaoTrademonitorGetRequest) GetPageNo() uint16 {
	page_no, found := req.Request.Params["page_no"]
	if found {
		return page_no.(uint16)
	}
	return 0
}

// 每页条数。（每页条数不超过50条，超过50或小于0均按50计）
func (req *FenxiaoTrademonitorGetRequest) SetPageSize(page_size uint16) {
	req.Request.Params["page_size"] = page_size
}

func (req *FenxiaoTrademonitorGetRequest) GetPageSize() uint16 {
	page_size, found := req.Request.Params["page_size"]
	if found {
		return page_size.(uint16)
	}
	return 0
}
