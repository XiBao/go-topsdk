package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type TbkItemConvertRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTbkItemConvertRequest() (req *TbkItemConvertRequest) {
	request := topsdk.Request{MethodName: "taobao.tbk.item.convert", Params: make(map[string]interface{}, 6)}
	req = &TbkItemConvertRequest{
		Request: &request,
	}
	return
}

// 返回值中需要的字段. 可选值 num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url字段间用 (,) 逗号分隔
func (req *TbkItemConvertRequest) SetFields(fields string) {
	req.Request.Params["fields"] = fields
}

func (req *TbkItemConvertRequest) GetFields() string {
	fields, found := req.Request.Params["fields"]
	if found {
		return fields.(string)
	}
	return ""
}

// 商品ID串，用,分割，从taobao.tbk.item.get接口获取num_iid字段，最大40个
func (req *TbkItemConvertRequest) SetNumIids(numIids string) {
	req.Request.Params["num_iids"] = numIids
}

func (req *TbkItemConvertRequest) GetNumIids() string {
	fields, found := req.Request.Params["num_iids"]
	if found {
		return fields.(string)
	}
	return ""
}

// mm_xxx_xxx_xxx的第三位,广告位ID，区分效果位置
func (req *TbkItemConvertRequest) SetAdzoneId(adzoneId uint64) {
	req.Request.Params["adzone_id"] = adzoneId
}

func (req *TbkItemConvertRequest) GetAdzoneId() uint64 {
	adzoneId, found := req.Request.Params["adzone_id"]
	if found {
		return adzoneId.(uint64)
	}
	return 0
}

// 链接形式：1：PC，2：无线，默认：１
func (req *TbkItemConvertRequest) SetPlatform(platform uint8) {
	req.Request.Params["platform"] = platform
}

func (req *TbkItemConvertRequest) GetPlatform() uint8 {
	platform, found := req.Request.Params["platform"]
	if found {
		return platform.(uint8)
	}
	return 1
}

// 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
func (req *TbkItemConvertRequest) SetUnid(unid string) {
	req.Request.Params["unid"] = unid
}

func (req *TbkItemConvertRequest) GetUnid() string {
	unid, found := req.Request.Params["unid"]
	if found {
		return unid.(string)
	}
	return ""
}

// 1表示商品转通用计划链接，其他值或不传表示转营销计划链接
func (req *TbkItemConvertRequest) SetDx(dx string) {
	req.Request.Params["dx"] = dx
}

func (req *TbkItemConvertRequest) GetDx() string {
	dx, found := req.Request.Params["dx"]
	if found {
		return dx.(string)
	}
	return ""
}
