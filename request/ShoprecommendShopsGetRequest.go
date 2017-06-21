package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ShoprecommendShopsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewShoprecommendShopsGetRequest() (req *ShoprecommendShopsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.shoprecommend.shops.get", Params: make(map[string]interface{}, 4)}
	req = &ShoprecommendShopsGetRequest{
		Request: &request,
	}
	return
}

// 传入卖家ID
func (req *ShoprecommendShopsGetRequest) SetSellerId(sellerId uint64) {
	req.Request.Params["seller_id"] = sellerId
}

func (req *ShoprecommendShopsGetRequest) GetSellerId() uint64 {
	sellerId, found := req.Request.Params["seller_id"]
	if found {
		return sellerId.(uint64)
	}
	return 0
}

// 请求类型，1：关联店铺推荐。其他值当非法值处理
func (req *ShoprecommendShopsGetRequest) SetRecommendType(recommendType uint8) {
	req.Request.Params["recommend_type"] = recommendType
}

func (req *ShoprecommendShopsGetRequest) GetRecommendType() uint8 {
	recommendType, found := req.Request.Params["recommend_type"]
	if found {
		return recommendType.(uint8)
	}
	return 0
}

// 请求个数，建议获取16个
func (req *ShoprecommendShopsGetRequest) SetCount(count uint) {
	req.Request.Params["count"] = count
}

func (req *ShoprecommendShopsGetRequest) GetCount() uint {
	count, found := req.Request.Params["count"]
	if found {
		return count.(uint)
	}
	return 0
}

// 额外参数
func (req *ShoprecommendShopsGetRequest) SetExt(ext string) {
	req.Request.Params["ext"] = ext
}

func (req *ShoprecommendShopsGetRequest) GetExt() string {
	ext, found := req.Request.Params["ext"]
	if found {
		return ext.(string)
	}
	return ""
}
