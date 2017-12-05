package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type TbkShopGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTbkShopGetRequest() (req *TbkShopGetRequest) {
	request := topsdk.Request{MethodName: "taobao.tbk.shop.get", Params: make(map[string]interface{}, 15)}
	req = &TbkShopGetRequest{
		Request: &request,
	}
	return
}

// 需返回的字段列表(user_id,shop_title,shop_type,seller_nick,pict_url,shop_url)
func (req *TbkShopGetRequest) SetFields(fields string) {
	req.Request.Params["fields"] = fields
}

func (req *TbkShopGetRequest) GetFields() string {
	fields, found := req.Request.Params["fields"]
	if found {
		return fields.(string)
	}
	return ""
}

// 查询词
func (req *TbkShopGetRequest) SetQ(q string) {
	req.Request.Params["q"] = q
}

func (req *TbkShopGetRequest) GetQ() string {
	q, found := req.Request.Params["q"]
	if found {
		return q.(string)
	}
	return ""
}

// 排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction） (commission_rate_des)
func (req *TbkShopGetRequest) SetSort(sort string) {
	req.Request.Params["sort"] = sort
}

func (req *TbkShopGetRequest) GetSort() string {
	sort, found := req.Request.Params["sort"]
	if found {
		return sort.(string)
	}
	return ""
}

// 是否商城的店铺，设置为true表示该是属于淘宝商城的店铺，设置为false或不设置表示不判断这个属性
func (req *TbkShopGetRequest) SetIsTmall(isTmall bool) {
	req.Request.Params["is_tmall"] = isTmall
}

func (req *TbkShopGetRequest) GetIsTmall() bool {
	isTmall, found := req.Request.Params["is_tmall"]
	if found {
		return isTmall.(bool)
	}
	return false
}

// 信用等级下限，1~20
func (req *TbkShopGetRequest) SetStartCredit(startCredit uint) {
	req.Request.Params["start_credit"] = startCredit
}

func (req *TbkShopGetRequest) GetStartCredit() uint {
	startCredit, found := req.Request.Params["start_credit"]
	if found {
		return startCredit.(uint)
	}
	return 0
}

// 信用等级上限，1~20
func (req *TbkShopGetRequest) SetEndCredit(endCredit uint) {
	req.Request.Params["end_credit"] = endCredit
}

func (req *TbkShopGetRequest) GetEndCredit() uint {
	endCredit, found := req.Request.Params["end_credit"]
	if found {
		return endCredit.(uint)
	}
	return 0
}

// 淘客佣金比率下限，1~10000
func (req *TbkShopGetRequest) SetStartCommissionRate(startCommissionRate uint) {
	req.Request.Params["start_commission_rate"] = startCommissionRate
}

func (req *TbkShopGetRequest) GetStartCommissionRate() uint {
	startCommissionRate, found := req.Request.Params["start_commission_rate"]
	if found {
		return startCommissionRate.(uint)
	}
	return 0
}

// 淘客佣金比率上限，1~10000
func (req *TbkShopGetRequest) SetEndCommissionRate(endCommissionRate uint) {
	req.Request.Params["end_commission_rate"] = endCommissionRate
}

func (req *TbkShopGetRequest) GetEndCommissionRate() uint {
	endCommissionRate, found := req.Request.Params["end_commission_rate"]
	if found {
		return endCommissionRate.(uint)
	}
	return 0
}

// 店铺商品总数下限
func (req *TbkShopGetRequest) SetStartTotalAction(startTotalAction uint) {
	req.Request.Params["start_total_action"] = startTotalAction
}

func (req *TbkShopGetRequest) GetStartTotalAction() uint {
	startTotalAction, found := req.Request.Params["start_total_action"]
	if found {
		return startTotalAction.(uint)
	}
	return 0
}

// 店铺商品总数上限
func (req *TbkShopGetRequest) SetEndTotalAction(endTotalAction uint) {
	req.Request.Params["end_total_action"] = endTotalAction
}

func (req *TbkShopGetRequest) GetEndTotalAction() uint {
	endTotalAction, found := req.Request.Params["end_total_action"]
	if found {
		return endTotalAction.(uint)
	}
	return 0
}

// 累计推广商品下限
func (req *TbkShopGetRequest) SetStartAuctionCount(startAuctionCount uint) {
	req.Request.Params["start_auction_count"] = startAuctionCount
}

func (req *TbkShopGetRequest) GetStartAuctionCount() uint {
	startAuctionCount, found := req.Request.Params["start_auction_count"]
	if found {
		return startAuctionCount.(uint)
	}
	return 0
}

// 累计推广商品上限
func (req *TbkShopGetRequest) SetEndAuctionCount(endAuctionCount uint) {
	req.Request.Params["end_auction_count"] = endAuctionCount
}

func (req *TbkShopGetRequest) GetEndAuctionCount() uint {
	endAuctionCount, found := req.Request.Params["end_auction_count"]
	if found {
		return endAuctionCount.(uint)
	}
	return 0
}

// 1：PC，2：无线，默认：1
func (req *TbkShopGetRequest) SetPlatform(platform uint) {
	req.Request.Params["platform"] = platform
}

func (req *TbkShopGetRequest) GetNumIids() uint {
	platform, found := req.Request.Params["platform"]
	if found {
		return platform.(uint)
	}
	return 1
}

// 页大小，默认20，1~100
func (req *TbkShopGetRequest) SetPageSize(pageSize uint) {
	req.Request.Params["page_size"] = pageSize
}

func (req *TbkShopGetRequest) GetPageSize() uint {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(uint)
	}
	return 20
}

// 第几页，默认：1（当后台类目和查询词均不指定的时候，最多出10000个结果，即page_no*page_size不能超过200；当指定类目或关键词的时候，则最多出100个结果）
func (req *TbkShopGetRequest) SetPageNo(pageNo uint) {
	req.Request.Params["page_no"] = pageNo
}

func (req *TbkShopGetRequest) GetPageNo() uint {
	pageNo, found := req.Request.Params["page_no"]
	if found {
		return pageNo.(uint)
	}
	return 1
}
