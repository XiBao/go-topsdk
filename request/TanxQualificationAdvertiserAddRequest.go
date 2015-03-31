package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/tanx"
)

type TanxQualificationAdvertiserAddRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTanxQualificationAdvertiserAddRequest() (req *TanxQualificationAdvertiserAddRequest) {
	request := topsdk.Request{MethodName: "taobao.tanx.qualification.advertiser.add", Params: make(map[string]interface{}, 4)}
	req = &TanxQualificationAdvertiserAddRequest{
		Request: &request,
	}
	return
}

// dsp用户memberId
func (req *TanxQualificationAdvertiserAddRequest) SetMemberId(memberId uint64) {
	req.Request.Params["member_id"] = memberId
}

func (req *TanxQualificationAdvertiserAddRequest) GetMemberId() uint64 {
	memberId, found := req.Request.Params["member_id"]
	if found {
		return memberId.(uint64)
	}
	return 0
}

// dsp用户验证token
func (req *TanxQualificationAdvertiserAddRequest) SetToken(token string) {
	req.Request.Params["token"] = token
}

func (req *TanxQualificationAdvertiserAddRequest) GetToken() string {
	token, found := req.Request.Params["token"]
	if found {
		return token.(string)
	}
	return ""
}

// 从1970年到当前时间的秒
func (req *TanxQualificationAdvertiserAddRequest) SetSignTime(signTime int64) {
	req.Request.Params["sign_time"] = signTime
}

func (req *TanxQualificationAdvertiserAddRequest) GetSignTime() int64 {
	signTime, found := req.Request.Params["sign_time"]
	if found {
		return signTime.(int64)
	}
	return 0
}

// 广告主对象
func (req *TanxQualificationAdvertiserAddRequest) SetAdvertisers(advertisers []*tanx.Advertiser) {
	js, _ := json.Marshal(advertisers)
	req.Request.Params["advertisers"] = string(js)
}

func (req *TanxQualificationAdvertiserAddRequest) GetAdvertisers() []*tanx.Advertiser {
	js, found := req.Request.Params["advertisers"]
	if found {
		advertisers := []*tanx.Advertiser{}
		json.Unmarshal([]byte(js.(string)), &advertisers)
		return advertisers
	}
	return nil
}
