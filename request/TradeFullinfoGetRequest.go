package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type TradeFullInfoGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTradeFullInfoGetRequest() (req *TradeFullInfoGetRequest) {
	request := topsdk.Request{MethodName: "taobao.trade.fullinfo.get", Params: make(map[string]interface{}, 2)}
	req = &TradeFullInfoGetRequest{
		Request: &request,
	}
	return
}

// 需要返回的字段。目前支持有：
// 1. Trade中可以指定返回的fields: seller_nick, buyer_nick, title, type, created, tid, seller_rate,seller_can_rate, buyer_rate,can_rate, status, payment, discount_fee, adjust_fee, post_fee, total_fee, pay_time, end_time, modified, consign_time, buyer_obtain_point_fee, point_fee, real_point_fee, received_payment, pic_path, num_iid, num, price, cod_fee, cod_status, shipping_type, receiver_name, receiver_state, receiver_city, receiver_district, receiver_address, receiver_zip, receiver_mobile, receiver_phone,seller_flag,alipay_id,alipay_no,is_lgtype,is_force_wlb,is_brand_sale,buyer_area,has_buyer_message, credit_card_fee, lg_aging_type, lg_aging, step_trade_status,step_paid_fee,mark_desc,has_yfx,yfx_fee,yfx_id,yfx_type,trade_source,send_time,is_daixiao,is_wt,is_part_consign
// 2. Order中可以指定返回fields：orders.title, orders.pic_path, orders.price, orders.num, orders.num_iid, orders.sku_id, orders.refund_status, orders.status, orders.oid, orders.total_fee, orders.payment, orders.discount_fee, orders.adjust_fee, orders.sku_properties_name, orders.item_meal_name, orders.buyer_rate, orders.seller_rate, orders.outer_iid, orders.outer_sku_id, orders.refund_id, orders.seller_type, orders.end_time,orders.order_from,orders.consign_time,orders.shipping_type,orders.logistics_company,orders.invoice_no,orders.is_daixiao
// 3. fields：orders（返回2中Order的所有内容）
// 4.fields:service_orders(返回service_order中所有内容)
func (req *TradeFullInfoGetRequest) SetFields(fields string) {
	req.Request.Params["fields"] = fields
}

func (req *TradeFullInfoGetRequest) GetFields() string {
	fields, found := req.Request.Params["fields"]
	if found {
		return fields.(string)
	}
	return ""
}

// 交易编号
func (req *TradeFullInfoGetRequest) SetTid(tid uint64) {
	req.Request.Params["tid"] = tid
}

func (req *TradeFullInfoGetRequest) GetTid() uint64 {
	tid, found := req.Request.Params["tid"]
	if found {
		return tid.(uint64)
	}
	return 0
}
