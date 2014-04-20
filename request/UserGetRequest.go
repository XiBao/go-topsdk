package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type UserGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewUserGetRequest() (req *UserGetRequest) {
	request := topsdk.Request{MethodName: "taobao.user.get", Params: make(map[string]interface{}, 2)}
	req = &UserGetRequest{
		Request: &request,
	}
	return
}

// 只返回user_id,nick,sex,seller_credit,type,has_more_pic,item_img_num,item_img_size,prop_img_num,prop_img_size,auto_repost,promoted_type,status,alipay_bind,consumer_protection,avatar,liangpin,sign_food_seller_promise,has_shop,is_lightning_consignment,has_sub_stock,is_golden_seller,vip_info,magazine_subscribe,vertical_market,online_gaming参数
func (req *UserGetRequest) SetFields(fields string) {
	req.Request.Params["fields"] = fields
}

func (req *UserGetRequest) GetFields() string {
	fields, found := req.Request.Params["fields"]
	if found {
		return fields.(string)
	}
	return ""
}

func (req *UserGetRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *UserGetRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}
