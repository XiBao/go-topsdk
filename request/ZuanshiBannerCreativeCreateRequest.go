package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerCreativeCreateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCreativeCreateRequest() (req *ZuanshiBannerCreativeCreateRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.creative.create", Params: make(map[string]interface{}, 5)}
	req = &ZuanshiBannerCreativeCreateRequest{
		Request: &request,
	}
	return
}

//创意名称
func (req *ZuanshiBannerCreativeCreateRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *ZuanshiBannerCreativeCreateRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

//是否需要pc转无线链接，true：是，false：否
func (req *ZuanshiBannerCreativeCreateRequest) SetIsTransToWifi(isTransToWifi bool) {
	req.Request.Params["is_trans_to_wifi"] = isTransToWifi
}

func (req *ZuanshiBannerCreativeCreateRequest) GetIsTransToWifi() bool {
	isTransToWifi, found := req.Request.Params["is_trans_to_wifi"]
	if found {
		return isTransToWifi.(bool)
	}
	return false
}

//类目ID，1,女装 2,彩妆，护肤 3,女鞋 4,男装 5,男鞋 6,箱包 7,食品，茶叶，烟酒 8,内衣 9,保健 10,饰品，服饰 11,运动 12,居家 13,童装，母婴 14,电器 15,数码 16,游戏 17,团购 18,其他
func (req *ZuanshiBannerCreativeCreateRequest) SetCatId(catId uint) {
	req.Request.Params["cat_id"] = catId
}

func (req *ZuanshiBannerCreativeCreateRequest) GetCatId() uint {
	catId, found := req.Request.Params["cat_id"]
	if found {
		return catId.(uint)
	}
	return 0
}

//点击链接
func (req *ZuanshiBannerCreativeCreateRequest) SetClickUrl(clickURL string) {
	req.Request.Params["click_url"] = clickURL
}

func (req *ZuanshiBannerCreativeCreateRequest) GetClickUrl() string {
	clickURL, found := req.Request.Params["click_url"]
	if found {
		return clickURL.(string)
	}
	return ""
}

// 图片二进制文件流
func (req *ZuanshiBannerCreativeCreateRequest) SetImage(image *topsdk.MultipartFile) {
	req.Request.Params["image"] = image
}

func (req *ZuanshiBannerCreativeCreateRequest) GetImage() *topsdk.MultipartFile {
	image, found := req.Request.Params["image"]
	if found {
		return image.(*topsdk.MultipartFile)
	}
	return nil
}
