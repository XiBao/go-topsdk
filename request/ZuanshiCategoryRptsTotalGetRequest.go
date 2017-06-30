package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiCategoryRptsTotalGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiCategoryRptsTotalGetRequest() (req *ZuanshiCategoryRptsTotalGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.category.rpts.total.get", Params: make(map[string]interface{}, 6)}
	req = &ZuanshiCategoryRptsTotalGetRequest{
		Request: &request,
	}
	return
}

//开始日期
func (req *ZuanshiCategoryRptsTotalGetRequest) SetStartTime(startTime string) {
	req.Request.Params["start_time"] = startTime
}

func (req *ZuanshiCategoryRptsTotalGetRequest) GetStartTime() string {
	startTime, found := req.Request.Params["start_time"]
	if found {
		return startTime.(string)
	}
	return ""
}

//结束日期
func (req *ZuanshiCategoryRptsTotalGetRequest) SetEndTime(endTime string) {
	req.Request.Params["end_time"] = endTime
}

func (req *ZuanshiCategoryRptsTotalGetRequest) GetEndTime() string {
	endTime, found := req.Request.Params["end_time"]
	if found {
		return endTime.(string)
	}
	return ""
}

// 计划类型。1：全店推广；4单品推广
func (req *ZuanshiCategoryRptsTotalGetRequest) SetCampaignModel(model uint8) {
	req.Request.Params["campaign_model"] = model
}

func (req *ZuanshiCategoryRptsTotalGetRequest) GetCampaignModel() uint8 {
	model, found := req.Request.Params["campaign_model"]
	if found {
		return model.(uint8)
	}
	return 0
}

// 效果周期，取值范围：3,7,15。分别表示效果转化周期-3天/7天/15天。
func (req *ZuanshiCategoryRptsTotalGetRequest) SetEffect(effect uint) {
	req.Request.Params["effect"] = effect
}

func (req *ZuanshiCategoryRptsTotalGetRequest) GetEffect() uint {
	effect, found := req.Request.Params["effect"]
	if found {
		return effect.(uint)
	}
	return 0
}

// 效果类型。“impression”：展现效果；“click”：点击效果
func (req *ZuanshiCategoryRptsTotalGetRequest) SetEffectType(effectType string) {
	req.Request.Params["effect_type"] = effectType
}

func (req *ZuanshiCategoryRptsTotalGetRequest) GetEffectType() string {
	effectType, found := req.Request.Params["effect_type"]
	if found {
		return effectType.(string)
	}
	return ""
}

// 类目id
func (req *ZuanshiCategoryRptsTotalGetRequest) SetShopMainCatId(shopMainCatId uint64) {
	req.Request.Params["shop_main_cat_id"] = shopMainCatId
}

func (req *ZuanshiCategoryRptsTotalGetRequest) GetShopMainCatId() uint64 {
	shopMainCatId, found := req.Request.Params["shop_main_cat_id"]
	if found {
		return shopMainCatId.(uint64)
	}
	return 0
}
