package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type TbkItemGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTbkItemGetRequest() (req *TbkItemGetRequest) {
	request := topsdk.Request{MethodName: "taobao.tbk.item.get", Params: make(map[string]interface{}, 14)}
	req = &TbkItemGetRequest{
		Request: &request,
	}
	return
}

// 返回值中需要的字段. 可选值 num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url字段间用 (,) 逗号分隔
func (req *TbkItemGetRequest) SetFields(fields string) {
	req.Request.Params["fields"] = fields
}

func (req *TbkItemGetRequest) GetFields() string {
	fields, found := req.Request.Params["fields"]
	if found {
		return fields.(string)
	}
	return ""
}

// 查询词
func (req *TbkItemGetRequest) SetQ(q string) {
	req.Request.Params["q"] = q
}

func (req *TbkItemGetRequest) GetQ() string {
	q, found := req.Request.Params["q"]
	if found {
		return q.(string)
	}
	return ""
}

// 后台类目ID，用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
func (req *TbkItemGetRequest) SetCat(cat string) {
	req.Request.Params["cat"] = cat
}

func (req *TbkItemGetRequest) GetCat() string {
	cat, found := req.Request.Params["cat"]
	if found {
		return cat.(string)
	}
	return ""
}

// 所在地
func (req *TbkItemGetRequest) SetItemloc(itemloc string) {
	req.Request.Params["itemloc"] = itemloc
}

func (req *TbkItemGetRequest) GetItemloc() string {
	itemloc, found := req.Request.Params["itemloc"]
	if found {
		return itemloc.(string)
	}
	return ""
}

// 排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi）
func (req *TbkItemGetRequest) SetSort(sort string) {
	req.Request.Params["sort"] = sort
}

func (req *TbkItemGetRequest) GetSort() string {
	sort, found := req.Request.Params["sort"]
	if found {
		return sort.(string)
	}
	return ""
}

// 是否商城商品，设置为true表示该商品是属于淘宝商城商品，设置为false或不设置表示不判断这个属性
func (req *TbkItemGetRequest) SetIsTmall(isTmall bool) {
	req.Request.Params["is_tmall"] = isTmall
}

func (req *TbkItemGetRequest) GetIsTmall() bool {
	isTmall, found := req.Request.Params["is_tmall"]
	if found {
		return isTmall.(bool)
	}
	return false
}

// 是否海外商品，设置为true表示该商品是属于海外商品，设置为false或不设置表示不判断这个属性
func (req *TbkItemGetRequest) SetIsOverseas(isOverseas bool) {
	req.Request.Params["is_overseas"] = isOverseas
}

func (req *TbkItemGetRequest) GetIsOverseas() bool {
	isOverseas, found := req.Request.Params["is_overseas"]
	if found {
		return isOverseas.(bool)
	}
	return false
}

// 折扣价范围下限，单位：元
func (req *TbkItemGetRequest) SetStartPrice(startPrice uint) {
	req.Request.Params["start_price"] = startPrice
}

func (req *TbkItemGetRequest) GetStartPrice() uint {
	startPrice, found := req.Request.Params["start_price"]
	if found {
		return startPrice.(uint)
	}
	return 0
}

// 折扣价范围上限，单位：元
func (req *TbkItemGetRequest) SetEndPrice(endPrice uint) {
	req.Request.Params["end_price"] = endPrice
}

func (req *TbkItemGetRequest) GetEndPrice() uint {
	endPrice, found := req.Request.Params["end_price"]
	if found {
		return endPrice.(uint)
	}
	return 0
}

// 淘客佣金比率上限，如：1234表示12.34%
func (req *TbkItemGetRequest) SetStartTkRate(startTkRate uint) {
	req.Request.Params["start_tk_rate"] = startTkRate
}

func (req *TbkItemGetRequest) GetStartTkRate() uint {
	startTkRate, found := req.Request.Params["start_tk_rate"]
	if found {
		return startTkRate.(uint)
	}
	return 0
}

// 淘客佣金比率下限，如：1234表示12.34%
func (req *TbkItemGetRequest) SetEndTkRate(endTkRate uint) {
	req.Request.Params["start_tk_rate"] = endTkRate
}

func (req *TbkItemGetRequest) GetEndTkRate() uint {
	endTkRate, found := req.Request.Params["end_tk_rate"]
	if found {
		return endTkRate.(uint)
	}
	return 0
}

// 第几页，默认：１
func (req *TbkItemGetRequest) SetPageNo(pageNo uint) {
	req.Request.Params["page_no"] = pageNo
}

func (req *TbkItemGetRequest) GetPageNo() uint {
	pageNo, found := req.Request.Params["page_no"]
	if found {
		return pageNo.(uint)
	}
	return 0
}

// 页大小，默认20，1~100
func (req *TbkItemGetRequest) SetPageSize(pageSize uint) {
	req.Request.Params["page_size"] = pageSize
}

func (req *TbkItemGetRequest) GetPageSize() uint {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(uint)
	}
	return 0
}

// 链接形式：1：PC，2：无线，默认：１
func (req *TbkItemGetRequest) SetPlatform(platform uint8) {
	req.Request.Params["platform"] = platform
}

func (req *TbkItemGetRequest) GetPlatform() uint8 {
	platform, found := req.Request.Params["platform"]
	if found {
		return platform.(uint8)
	}
	return 1
}
