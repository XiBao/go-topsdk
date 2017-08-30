package request

import (
	"github.com/XiBao/taobao"
	"github.com/XiBao/topsdk"
)

type TopSdkFeedbackUploadRequest struct {
	Request *topsdk.Request
	ApiKey  *taobao.ApiKey
}

// create new request
func NewTopSdkFeedbackUploadRequest(apiKey *taobao.ApiKey) (req *TopSdkFeedbackUploadRequest) {
	request := topsdk.Request{MethodName: "taobao.top.sdk.feedback.upload", Params: make(map[string]interface{})}
	req = &TopSdkFeedbackUploadRequest{
		Request: &request,
		ApiKey:  apiKey,
	}
	return
}

func (req *TopSdkFeedbackUploadRequest) SetUploadType(uploadType string) {
	req.Request.Params["type"] = uploadType
}

func (req *TopSdkFeedbackUploadRequest) GetUploadType() string {
	uploadType, found := req.Request.Params["type"]
	if found {
		return uploadType.(string)
	}
	return ""
}

func (req *TopSdkFeedbackUploadRequest) SetContext(content string) {
	req.Request.Params["content"] = content
}

func (req *TopSdkFeedbackUploadRequest) GetContext() string {
	content, found := req.Request.Params["content"]
	if found {
		return content.(string)
	}
	return ""
}
