package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type FenxiaoOrdersGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewFenxiaoOrdersGetRequest() (req *FenxiaoOrdersGetRequest) {
	request := topsdk.Request{MethodName: "taobao.fenxiao.orders.get", Params: make(map[string]interface{}, 9)}
	req = &FenxiaoOrdersGetRequest{
		Request: &request,
	}
	return
}

// 结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。。
func (req *FenxiaoOrdersGetRequest) SetEndCreated(endCreated string) {
	req.Request.Params["end_created"] = endCreated
}

func (req *FenxiaoOrdersGetRequest) GetEndCreated() string {
	endCreated, found := req.Request.Params["end_created"]
	if found {
		return endCreated.(string)
	}
	return ""
}

// 起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。
func (req *FenxiaoOrdersGetRequest) SetStartCreated(startCreated string) {
	req.Request.Params["start_created"] = startCreated
}

func (req *FenxiaoOrdersGetRequest) GetStartCreated() string {
	startCreated, found := req.Request.Params["start_created"]
	if found {
		return startCreated.(string)
	}
	return ""
}

/* 多个字段用","分隔。
fields
如果为空：返回所有采购单对象(purchase_orders)字段。
如果不为空：返回指定采购单对象(purchase_orders)字段。

例1：
sub_purchase_orders.tc_order_id 表示只返回tc_order_id
例2：
sub_purchase_orders表示只返回子采购单列表
*/
func (req *FenxiaoOrdersGetRequest) SetFields(fields string) {
	req.Request.Params["fields"] = fields
}

func (req *FenxiaoOrdersGetRequest) GetFields() string {
	fields, found := req.Request.Params["fields"]
	if found {
		return fields.(string)
	}
	return ""
}

// 采购单下游买家订单id
func (req *FenxiaoOrdersGetRequest) SetTcOrderId(tcOrderId uint64) {
	req.Request.Params["tc_order_id"] = tcOrderId
}

func (req *FenxiaoOrdersGetRequest) GetTcOrderId() uint64 {
	tcOrderId, found := req.Request.Params["tc_order_id"]
	if found {
		return tcOrderId.(uint64)
	}
	return 0
}

// 交易状态，不传默认查询所有采购单根据身份选择自身状态可选值:<br> *供应商：<br> WAIT_SELLER_SEND_GOODS(等待发货)<br> WAIT_SELLER_CONFIRM_PAY(待确认收款)<br> WAIT_BUYER_PAY(等待付款)<br> WAIT_BUYER_CONFIRM_GOODS(已发货)<br> TRADE_REFUNDING(退款中)<br> TRADE_FINISHED(采购成功)<br> TRADE_CLOSED(已关闭)<br> *分销商：<br> WAIT_BUYER_PAY(等待付款)<br> WAIT_BUYER_CONFIRM_GOODS(待收货确认)<br> TRADE_FOR_PAY(已付款)<br> TRADE_REFUNDING(退款中)<br> TRADE_FINISHED(采购成功)<br> TRADE_CLOSED(已关闭)<br>
func (req *FenxiaoOrdersGetRequest) SetStatus(status string) {
	req.Request.Params["status"] = status
}

func (req *FenxiaoOrdersGetRequest) GetStatus() string {
	status, found := req.Request.Params["status"]
	if found {
		return status.(string)
	}
	return ""
}

// 可选值：trade_time_type(采购单按照成交时间范围查询),update_time_type(采购单按照更新时间范围查询)
func (req *FenxiaoOrdersGetRequest) SetTimeType(timeType string) {
	req.Request.Params["time_type"] = timeType
}

func (req *FenxiaoOrdersGetRequest) GetTimeType() string {
	timeType, found := req.Request.Params["time_type"]
	if found {
		return timeType.(string)
	}
	return ""
}

// 采购单编号或分销流水号，若其它参数没传，则此参数必传。
func (req *FenxiaoOrdersGetRequest) SetPurchaseOrderId(purchaseOrderId uint64) {
	req.Request.Params["purchase_order_id"] = purchaseOrderId
}

func (req *FenxiaoOrdersGetRequest) GetPurchaseOrderId() uint64 {
	purchaseOrderId, found := req.Request.Params["purchase_order_id"]
	if found {
		return purchaseOrderId.(uint64)
	}
	return 0
}

// 页码。（大于0的整数。默认为1）
func (req *FenxiaoOrdersGetRequest) SetPageNo(page_no uint16) {
	req.Request.Params["page_no"] = page_no
}

func (req *FenxiaoOrdersGetRequest) GetPageNo() uint16 {
	page_no, found := req.Request.Params["page_no"]
	if found {
		return page_no.(uint16)
	}
	return 0
}

// 每页条数。（每页条数不超过50条）
func (req *FenxiaoOrdersGetRequest) SetPageSize(page_size uint16) {
	req.Request.Params["page_size"] = page_size
}

func (req *FenxiaoOrdersGetRequest) GetPageSize() uint16 {
	page_size, found := req.Request.Params["page_size"]
	if found {
		return page_size.(uint16)
	}
	return 0
}
