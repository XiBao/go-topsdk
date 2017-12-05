package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type TbkDgItemCouponGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTbkDgItemCouponGetRequest() (req *TbkDgItemCouponGetRequest) {
	request := topsdk.Request{MethodName: "taobao.tbk.dg.item.coupon.get", Params: make(map[string]interface{}, 6)}
	req = &TbkDgItemCouponGetRequest{
		Request: &request,
	}
	return
}

// mm_xxx_xxx_xxx的第三位
func (req *TbkDgItemCouponGetRequest) SetAdzoneId(adzoneId uint64) {
	req.Request.Params["adzone_id"] = adzoneId
}

func (req *TbkDgItemCouponGetRequest) GetAdzoneId() uint64 {
	adzoneId, found := req.Request.Params["adzone_id"]
	if found {
		return adzoneId.(uint64)
	}
	return 0
}

// 1：PC，2：无线，默认：1
func (req *TbkDgItemCouponGetRequest) SetPlatform(platform uint) {
	req.Request.Params["platform"] = platform
}

func (req *TbkDgItemCouponGetRequest) GetNumIids() uint {
	platform, found := req.Request.Params["platform"]
	if found {
		return platform.(uint)
	}
	return 1
}

// 后台类目ID，用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
func (req *TbkDgItemCouponGetRequest) SetCat(cat string) {
	req.Request.Params["cat"] = cat
}

func (req *TbkDgItemCouponGetRequest) GetCat() string {
	cat, found := req.Request.Params["cat"]
	if found {
		return cat.(string)
	}
	return ""
}

// 页大小，默认20，1~100
func (req *TbkDgItemCouponGetRequest) SetPageSize(pageSize uint) {
	req.Request.Params["page_size"] = pageSize
}

func (req *TbkDgItemCouponGetRequest) GetPageSize() uint {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(uint)
	}
	return 20
}

// 第几页，默认：1（当后台类目和查询词均不指定的时候，最多出10000个结果，即page_no*page_size不能超过200；当指定类目或关键词的时候，则最多出100个结果）
func (req *TbkDgItemCouponGetRequest) SetPageNo(pageNo uint) {
	req.Request.Params["page_no"] = pageNo
}

func (req *TbkDgItemCouponGetRequest) GetPageNo() uint {
	pageNo, found := req.Request.Params["page_no"]
	if found {
		return pageNo.(uint)
	}
	return 1
}

// 查询词
func (req *TbkDgItemCouponGetRequest) SetQ(q string) {
	req.Request.Params["q"] = q
}

func (req *TbkDgItemCouponGetRequest) GetQ() string {
	q, found := req.Request.Params["q"]
	if found {
		return q.(string)
	}
	return ""
}
