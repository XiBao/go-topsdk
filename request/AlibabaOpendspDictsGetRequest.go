package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type AlibabaOpendspDictsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspDictsGetRequest() (req *AlibabaOpendspDictsGetRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.dicts.get", Params: make(map[string]interface{}, 1)}
	req = &AlibabaOpendspDictsGetRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspDictsGetRequest) SetDspCustId(keyId string) {
	req.Request.Params["key_id"] = keyId
}

func (req *AlibabaOpendspDictsGetRequest) GetDspCustId() string {
	keyId, found := req.Request.Params["key_id"]
	if found {
		return keyId.(string)
	}
	return ""
}
