package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type TbkItemInfoGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTbkItemInfoGetRequest() (req *TbkItemInfoGetRequest) {
	request := topsdk.Request{MethodName: "taobao.tbk.item.info.get", Params: make(map[string]interface{}, 3)}
	req = &TbkItemInfoGetRequest{
		Request: &request,
	}
	return
}

// 返回值中需要的字段. 可选值 num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url字段间用 (,) 逗号分隔
func (req *TbkItemInfoGetRequest) SetFields(fields string) {
	req.Request.Params["fields"] = fields
}

func (req *TbkItemInfoGetRequest) GetFields() string {
	fields, found := req.Request.Params["fields"]
	if found {
		return fields.(string)
	}
	return ""
}

// 商品ID串，用,分割，从taobao.tbk.item.get接口获取num_iid字段，最大40个
func (req *TbkItemInfoGetRequest) SetNumIids(numIids string) {
	req.Request.Params["num_iids"] = numIids
}

func (req *TbkItemInfoGetRequest) GetNumIids() string {
	fields, found := req.Request.Params["num_iids"]
	if found {
		return fields.(string)
	}
	return ""
}

// 链接形式：1：PC，2：无线，默认：１
func (req *TbkItemInfoGetRequest) SetPlatform(platform uint8) {
	req.Request.Params["platform"] = platform
}

func (req *TbkItemInfoGetRequest) GetPlatform() uint8 {
	platform, found := req.Request.Params["platform"]
	if found {
		return platform.(uint8)
	}
	return 1
}
