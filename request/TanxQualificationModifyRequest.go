package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/tanx"
)

type TanxQualificationModifyRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTanxQualificationModifyRequest() (req *TanxQualificationModifyRequest) {
	request := topsdk.Request{MethodName: "taobao.tanx.qualification.modify", Params: make(map[string]interface{}, 4)}
	req = &TanxQualificationModifyRequest{
		Request: &request,
	}
	return
}

// dsp用户memberId
func (req *TanxQualificationModifyRequest) SetMemberId(memberId uint64) {
	req.Request.Params["member_id"] = memberId
}

func (req *TanxQualificationModifyRequest) GetMemberId() uint64 {
	memberId, found := req.Request.Params["member_id"]
	if found {
		return memberId.(uint64)
	}
	return 0
}

// dsp用户验证token
func (req *TanxQualificationModifyRequest) SetToken(token string) {
	req.Request.Params["token"] = token
}

func (req *TanxQualificationModifyRequest) GetToken() string {
	token, found := req.Request.Params["token"]
	if found {
		return token.(string)
	}
	return ""
}

// 从1970年到当前时间的秒
func (req *TanxQualificationModifyRequest) SetSignTime(signTime uint64) {
	req.Request.Params["sign_time"] = signTime
}

func (req *TanxQualificationModifyRequest) GetSignTime() uint64 {
	signTime, found := req.Request.Params["sign_time"]
	if found {
		return signTime.(uint64)
	}
	return 0
}

// dsp客户新增资质dto
func (req *TanxQualificationModifyRequest) SetQualifications(qualifications []*tanx.Qualification) {
	js, _ := json.Marshal(qualifications)
	req.Request.Params["qualifications"] = string(js)
}

func (req *TanxQualificationModifyRequest) GetQualifications() []*tanx.Qualification {
	js, found := req.Request.Params["qualifications"]
	if found {
		qualifications := []*tanx.Qualification{}
		json.Unmarshal([]byte(js.(string)), &qualifications)
		return qualifications
	}
	return nil
}
