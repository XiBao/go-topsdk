package TopRequest

import (
	"github.com/XiBao/topsdk"
	"time"
)

type PictureCategoryGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewPictureCategoryGetRequest() (req *PictureCategoryGetRequest) {
	request := topsdk.Request{MethodName: "taobao.picture.category.get", Params: make(map[string]interface{})}
	req = &PictureCategoryGetRequest{
		Request: &request,
	}
	return
}

//图片分类ID
func (req *PictureCategoryGetRequest) SetPicCateId(picCateId uint64) {
	req.Request.Params["picture_category_id"] = picCateId
}

func (req *PictureCategoryGetRequest) GetPicCateId() uint64 {
	picCateId, found := req.Request.Params["picture_category_id"]
	if found {
		return picCateId.(uint64)
	}
	return 0
}

//图片分类名称
func (req *PictureCategoryGetRequest) SetPicCateName(picCateName string) {
	req.Request.Params["picture_category_name"] = picCateName
}

func (req *PictureCategoryGetRequest) GetPicCateName() string {
	picCateName, found := req.Request.Params["picture_category_name"]
	if found {
		return picCateName.(string)
	}
	return ""
}

//分类类型,fixed代表店铺装修分类类别，auction代表宝贝分类类别，user-define代表用户自定义分类类别
func (req *PictureCategoryGetRequest) SetType(t string) {
	req.Request.Params["type"] = t
}

func (req *PictureCategoryGetRequest) GetType() string {
	t, found := req.Request.Params["type"]
	if found {
		return t.(string)
	}
	return ""
}

// 图片分类的父分类,一级分类的parent_id为0,二级分类的则为其父分类的picture_category_id
func (req *PictureCategoryGetRequest) SetParentId(parentId uint64) {
	req.Request.Params["parent_id"] = parentId
}

func (req *PictureCategoryGetRequest) GetParentId() uint64 {
	parentId, found := req.Request.Params["parent_id"]
	if found {
		return parentId.(uint64)
	}
	return 0
}

// 图片分类的父分类,一级分类的parent_id为0,二级分类的则为其父分类的picture_category_id
func (req *PictureCategoryGetRequest) SetModifiedTime(modifiedTime time.Time) {
	req.Request.Params["modified_time"] = modifiedTime
}

func (req *PictureCategoryGetRequest) GetModifiedTime() time.Time {
	modifiedTime, found := req.Request.Params["modified_time"]
	if found {
		return modifiedTime.(time.Time)
	}
	return time.Time{}
}
