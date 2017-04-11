package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type PictureDeleteRequest struct {
	Request *topsdk.Request
}

// create new request
func NewPictureDeleteRequest() (req *PictureDeleteRequest) {
	request := topsdk.Request{MethodName: "taobao.picture.delete", Params: make(map[string]interface{}, 1)}
	req = &PictureDeleteRequest{
		Request: &request,
	}
	return
}

func (req *PictureDeleteRequest) SetPictureIds(pictureIds string) {
	req.Request.Params["picture_ids"] = pictureIds
}

func (req *PictureDeleteRequest) GetPictureIds() string {
	pictureIds, found := req.Request.Params["picture_ids"]
	if found {
		return pictureIds.(string)
	}
	return ""
}
