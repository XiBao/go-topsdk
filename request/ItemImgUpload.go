package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ItemImgUploadRequest struct {
	Request *topsdk.Request
}

// create new request
func NewItemImgUploadRequest() (req *ItemImgUploadRequest) {
	request := topsdk.Request{MethodName: "taobao.item.img.upload", Params: make(map[string]interface{}, 2)}
	req = &ItemImgUploadRequest{
		Request: &request,
	}
	return
}

func (req *ItemImgUploadRequest) SetId(id uint64) {
	req.Request.Params["id"] = id
}

func (req *ItemImgUploadRequest) GetId() uint64 {
	id, found := req.Request.Params["id"]
	if found {
		return id.(uint64)
	}
	return 0
}

func (req *ItemImgUploadRequest) SetNumIid(numiid uint64) {
	req.Request.Params["num_iid"] = numiid
}

func (req *ItemImgUploadRequest) GetNumIid() uint64 {
	numiid, found := req.Request.Params["num_iid"]
	if found {
		return numiid.(uint64)
	}
	return 0
}

func (req *ItemImgUploadRequest) SetPosition(position uint64) {
	req.Request.Params["position"] = position
}

func (req *ItemImgUploadRequest) GetPosition() uint64 {
	position, found := req.Request.Params["position"]
	if found {
		return position.(uint64)
	}
	return 0
}

// 图片二进制文件流
// 商品图片内容类型:JPG,GIF;最大:3M 。支持的文件类型：gif,jpg,jpeg,png 支持的文件类型：gif,jpg,jpeg,png,bmp
func (req *ItemImgUploadRequest) SetImage(image *topsdk.MultipartFile) {
	req.Request.Params["image"] = image
}

func (req *ItemImgUploadRequest) GetImage() *topsdk.MultipartFile {
	image, found := req.Request.Params["image"]
	if found {
		return image.(*topsdk.MultipartFile)
	}
	return nil
}

func (req *ItemImgUploadRequest) SetIsMajor(isMajor bool) {
	req.Request.Params["is_major"] = isMajor
}

func (req *ItemImgUploadRequest) GetIsMajor() bool {
	isMajor, found := req.Request.Params["is_major"]
	if found {
		return isMajor.(bool)
	}
	return false
}
