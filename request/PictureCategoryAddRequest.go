package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type PictureCategoryAddRequest struct {
	Request *topsdk.Request
}

// create new request
func NewPictureCategoryAddRequest() (req *PictureCategoryAddRequest) {
	request := topsdk.Request{MethodName: "taobao.picture.category.add", Params: make(map[string]interface{}, 1)}
	req = &PictureCategoryAddRequest{
		Request: &request,
	}
	return
}

//图片分类名称
func (req *PictureCategoryAddRequest) SetPicCateName(picCateName string) {
	req.Request.Params["picture_category_name"] = picCateName
}

func (req *PictureCategoryAddRequest) GetPicCateName() string {
	picCateName, found := req.Request.Params["picture_category_name"]
	if found {
		return picCateName.(string)
	}
	return ""
}

// 图片分类的父分类,一级分类的parent_id为0,二级分类的则为其父分类的picture_category_id
func (req *PictureCategoryAddRequest) SetParentId(parentId uint64) {
	req.Request.Params["parent_id"] = parentId
}

func (req *PictureCategoryAddRequest) GetParentId() uint64 {
	parentId, found := req.Request.Params["parent_id"]
	if found {
		return parentId.(uint64)
	}
	return 0
}
