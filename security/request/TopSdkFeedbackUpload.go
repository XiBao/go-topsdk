package request

import (
	"fmt"
	"github.com/XiBao/topsdk"
)

func TopSdkFeedbackUpload(req *TopSdkFeedbackUploadRequest) (uploadInterval uint64, err error) {
	c := topsdk.NewClient(req.ApiKey.Key, req.ApiKey.Secret)
	result, err := c.Execute(req.Request)
	if err != nil {
		return 0, err
	}
	v, ok := result.(uint64)
	if !ok {
		return 0, fmt.Errorf("invalid result: %v", result)
	}
	return v, nil
}
