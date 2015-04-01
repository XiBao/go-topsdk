package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/tanx"
)

type TanxQualificationAddRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTanxQualificationAddRequest() (req *TanxQualificationAddRequest) {
	request := topsdk.Request{MethodName: "taobao.tanx.qualification.add", Params: make(map[string]interface{}, 4)}
	req = &TanxQualificationAddRequest{
		Request: &request,
	}
	return
}

// dsp用户memberId
func (req *TanxQualificationAddRequest) SetMemberId(memberId uint64) {
	req.Request.Params["member_id"] = memberId
}

func (req *TanxQualificationAddRequest) GetMemberId() uint64 {
	memberId, found := req.Request.Params["member_id"]
	if found {
		return memberId.(uint64)
	}
	return 0
}

// dsp用户验证token
func (req *TanxQualificationAddRequest) SetToken(token string) {
	req.Request.Params["token"] = token
}

func (req *TanxQualificationAddRequest) GetToken() string {
	token, found := req.Request.Params["token"]
	if found {
		return token.(string)
	}
	return ""
}

// 从1970年到当前时间的秒
func (req *TanxQualificationAddRequest) SetSignTime(signTime int64) {
	req.Request.Params["sign_time"] = signTime
}

func (req *TanxQualificationAddRequest) GetSignTime() int64 {
	signTime, found := req.Request.Params["sign_time"]
	if found {
		return signTime.(int64)
	}
	return 0
}

// dsp客户新增资质dto
func (req *TanxQualificationAddRequest) SetQualifications(qualifications []*tanx.Qualification) {
	js, _ := json.Marshal(qualifications)
	req.Request.Params["qualifications"] = string(js)
}

func (req *TanxQualificationAddRequest) GetQualifications() []*tanx.Qualification {
	js, found := req.Request.Params["qualifications"]
	if found {
		qualifications := []*tanx.Qualification{}
		json.Unmarshal([]byte(js.(string)), &qualifications)
		return qualifications
	}
	return nil
}
