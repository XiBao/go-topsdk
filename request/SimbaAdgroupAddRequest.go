package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaAdgroupAddRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaAdgroupAddRequest() (req *SimbaAdgroupAddRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.adgroup.add", Params:make(map[string]interface{}, 6)}
    req = &SimbaAdgroupAddRequest {
        Request: &request,
    }
    return
}

// 推广计划Id
func (req *SimbaAdgroupAddRequest) SetCampaignId(campaign_id uint64) {
    req.Request.Params["campaign_id"] = campaign_id
}

func (req *SimbaAdgroupAddRequest) GetCampaignId() uint64 {
    campaign_id, found := req.Request.Params["campaign_id"]
    if found {
        return campaign_id.(uint64)
    }
    return 0
}

//推广组默认出价，单位为分，不能小于5 不能大于日最高限额
func (req *SimbaAdgroupAddRequest) SetDefaultPrice(default_price uint32) {
    req.Request.Params["default_price"] = default_price
}

func (req *SimbaAdgroupAddRequest) GetDefaultPrice() uint32 {
    default_price, found := req.Request.Params["default_price"]
    if found {
        return default_price.(uint32)
    }
    return 0
}

//主人昵称
func (req *SimbaAdgroupAddRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaAdgroupAddRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}

//创意图片地址，必须是商品的图片之一
func (req *SimbaAdgroupAddRequest) SetImgUrl(img_url string) {
    req.Request.Params["img_url"] = img_url
}

func (req *SimbaAdgroupAddRequest) GetImgUrl() string {
    img_url, found := req.Request.Params["img_url"]
    if found {
        return img_url.(string)
    }
    return ""
}

// 商品Id
func (req *SimbaAdgroupAddRequest) SetItemId(item_id uint64) {
    req.Request.Params["item_id"] = item_id
}

func (req *SimbaAdgroupAddRequest) GetItemId() uint64 {
    item_id, found := req.Request.Params["item_id"]
    if found {
        return item_id.(uint64)
    }
    return 0
}

//创意标题，最多20个汉字
func (req *SimbaAdgroupAddRequest) SetTitle(title string) {
    req.Request.Params["title"] = title
}

func (req *SimbaAdgroupAddRequest) GetTitle() string {
    title, found := req.Request.Params["title"]
    if found {
        return title.(string)
    }
    return ""
}