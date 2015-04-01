package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type TanxQualificationSolidFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTanxQualificationSolidFindRequest() (req *TanxQualificationSolidFindRequest) {
	request := topsdk.Request{MethodName: "taobao.tanx.qualification.solid.find", Params: make(map[string]interface{}, 7)}
	req = &TanxQualificationSolidFindRequest{
		Request: &request,
	}
	return
}

// dsp用户memberId
func (req *TanxQualificationSolidFindRequest) SetMemberId(memberId uint64) {
	req.Request.Params["member_id"] = memberId
}

func (req *TanxQualificationSolidFindRequest) GetMemberId() uint64 {
	memberId, found := req.Request.Params["member_id"]
	if found {
		return memberId.(uint64)
	}
	return 0
}

// dsp用户验证token
func (req *TanxQualificationSolidFindRequest) SetToken(token string) {
	req.Request.Params["token"] = token
}

func (req *TanxQualificationSolidFindRequest) GetToken() string {
	token, found := req.Request.Params["token"]
	if found {
		return token.(string)
	}
	return ""
}

// 从1970年到当前时间的秒
func (req *TanxQualificationSolidFindRequest) SetSignTime(signTime int64) {
	req.Request.Params["sign_time"] = signTime
}

func (req *TanxQualificationSolidFindRequest) GetSignTime() int64 {
	signTime, found := req.Request.Params["sign_time"]
	if found {
		return signTime.(int64)
	}
	return 0
}

// 资质元素id列表
func (req *TanxQualificationSolidFindRequest) SetElementIds(elemetnIds string) {
	req.Request.Params["element_ids"] = elemetnIds
}

func (req *TanxQualificationSolidFindRequest) GetElementIds() string {
	elemetnIds, found := req.Request.Params["element_ids"]
	if found {
		return elemetnIds.(string)
	}
	return ""
}

// 广告主id
func (req *TanxQualificationSolidFindRequest) SetAdvertiserId(advertiserId uint64) {
	req.Request.Params["advertiser_id"] = advertiserId
}

func (req *TanxQualificationSolidFindRequest) GetAdvertiserId() uint64 {
	advertiserId, found := req.Request.Params["advertiser_id"]
	if found {
		return advertiserId.(uint64)
	}
	return 0
}

// 查询起始页
func (req *TanxQualificationSolidFindRequest) SetPage(page uint) {
	req.Request.Params["page"] = page
}

func (req *TanxQualificationSolidFindRequest) GetPage() uint {
	page, found := req.Request.Params["page"]
	if found {
		return page.(uint)
	}
	return 0
}

// 分页大小
func (req *TanxQualificationSolidFindRequest) SetPageSize(pageSize uint) {
	req.Request.Params["page_size"] = pageSize
}

func (req *TanxQualificationSolidFindRequest) GetSize() uint {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(uint)
	}
	return 0
}
