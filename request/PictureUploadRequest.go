package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type PictureUploadRequest struct {
	Request *topsdk.Request
}

// create new request
func NewPictureUploadRequest() (req *PictureUploadRequest) {
	request := topsdk.Request{MethodName: "taobao.picture.upload", Params: make(map[string]interface{}, 3)}
	req = &PictureUploadRequest{
		Request: &request,
	}
	return
}

//图片分类ID
func (req *PictureUploadRequest) SetPicCateId(picCateId uint64) {
	req.Request.Params["picture_category_id"] = picCateId
}

func (req *PictureUploadRequest) GetPicCateId() uint64 {
	picCateId, found := req.Request.Params["picture_category_id"]
	if found {
		return picCateId.(uint64)
	}
	return 0
}

// 图片二进制文件流
func (req *PictureUploadRequest) SetImg(img []byte) {
	req.Request.Params["img"] = img
}

func (req *PictureUploadRequest) GetImg() []byte {
	img, found := req.Request.Params["img"]
	if found {
		return img.([]byte)
	}
	return nil
}

// 包括后缀名的图片标题
func (req *PictureUploadRequest) SetImgInputTitle(imgInputTitle string) {
	req.Request.Params["image_input_title"] = imgInputTitle
}

func (req *PictureUploadRequest) GetImgInputTitle() string {
	imgInputTitle, found := req.Request.Params["image_input_title"]
	if found {
		return imgInputTitle.(string)
	}
	return ""
}

// 图片标题
func (req *PictureUploadRequest) SetTitle(title string) {
	req.Request.Params["title"] = title
}

// 图片上传的来源
func (req *PictureUploadRequest) GetTitle() string {
	title, found := req.Request.Params["title"]
	if found {
		return title.(string)
	}
	return "client:computer"
}

func (req *PictureUploadRequest) SetClientType(clientType string) {
	req.Request.Params["client_type"] = clientType
}

func (req *PictureUploadRequest) GetClientType() string {
	clientType, found := req.Request.Params["client_type"]
	if found {
		return clientType.(string)
	}
	return ""
}

func (req *PictureUploadRequest) SetIsHttps(isHttps bool) {
	req.Request.Params["is_https"] = isHttps
}

func (req *PictureUploadRequest) GetIsHttps() bool {
	isHttps, found := req.Request.Params["is_https"]
	if found {
		return isHttps.(bool)
	}
	return false
}
