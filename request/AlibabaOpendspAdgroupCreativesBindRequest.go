package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type AlibabaOpendspAdgroupCreativesBindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspAdgroupCreativesBindRequest() (req *AlibabaOpendspAdgroupCreativesBindRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.adgroup.creatives.bind", Params: make(map[string]interface{}, 3)}
	req = &AlibabaOpendspAdgroupCreativesBindRequest{
		Request: &request,
	}
	return
}

// DSP系统中推广主id
func (req *AlibabaOpendspAdgroupCreativesBindRequest) SetDspCustId(dspCustId string) {
	req.Request.Params["dsp_cust_id"] = dspCustId
}

func (req *AlibabaOpendspAdgroupCreativesBindRequest) GetAdgroupList() string {
	dspCustId, found := req.Request.Params["dsp_cust_id"]
	if found {
		return dspCustId.(string)
	}
	return ""
}

// 推广单元ID
func (req *AlibabaOpendspAdgroupCreativesBindRequest) SetAdgroupId(adgroupId uint64) {
	req.Request.Params["adgroup_id"] = adgroupId
}

func (req *AlibabaOpendspAdgroupCreativesBindRequest) GetAdgroupId() uint64 {
	adgroupId, found := req.Request.Params["adgroup_id"]
	if found {
		return adgroupId.(uint64)
	}
	return 0
}

// 创意ID列表
func (req *AlibabaOpendspAdgroupCreativesBindRequest) SetCreativeIdList(ids string) {
	req.Request.Params["creative_id_list"] = ids
}

func (req *AlibabaOpendspAdgroupCreativesBindRequest) GetCreativeIdList() string {
	ids, found := req.Request.Params["creative_id_list"]
	if found {
		return ids.(string)
	}
	return ""
}
