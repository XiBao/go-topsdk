package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type PicturesCountRequest struct {
	Request *topsdk.Request
}

// create new request
func NewPicturesCountRequest() (req *PicturesCountRequest) {
	request := topsdk.Request{MethodName: "taobao.picture.pictures.count", Params: make(map[string]interface{})}
	req = &PicturesCountRequest{
		Request: &request,
	}
	return
}

func (req *PicturesCountRequest) SetStartDate(startDate string) {
	req.Request.Params["start_date"] = startDate
}

func (req *PicturesCountRequest) GetStartDate() string {
	startDate, found := req.Request.Params["start_date"]
	if found {
		return startDate.(string)
	}
	return ""
}

//图片分类ID
func (req *PicturesCountRequest) SetPicCateId(picCateId uint64) {
	req.Request.Params["picture_category_id"] = picCateId
}

func (req *PicturesCountRequest) GetPicCateId() uint64 {
	picCateId, found := req.Request.Params["picture_category_id"]
	if found {
		return picCateId.(uint64)
	}
	return 0
}

func (req *PicturesCountRequest) SetTitle(title string) {
	req.Request.Params["title"] = title
}

func (req *PicturesCountRequest) GetTitle() string {
	title, found := req.Request.Params["title"]
	if found {
		return title.(string)
	}
	return ""
}

func (req *PicturesCountRequest) SetEndDate(endDate string) {
	req.Request.Params["end_date"] = endDate
}

func (req *PicturesCountRequest) GetEndDate() string {
	endDate, found := req.Request.Params["end_date"]
	if found {
		return endDate.(string)
	}
	return ""
}

func (req *PicturesCountRequest) SetStartModifiedDate(startModifiedDate string) {
	req.Request.Params["start_modified_date"] = startModifiedDate
}

func (req *PicturesCountRequest) GetStartModifiedDate() string {
	startModifiedDate, found := req.Request.Params["start_modified_date"]
	if found {
		return startModifiedDate.(string)
	}
	return ""
}

func (req *PicturesCountRequest) SetDeleted(deleted string) {
	req.Request.Params["deleted"] = deleted
}

func (req *PicturesCountRequest) GetDeleted() string {
	deleted, found := req.Request.Params["deleted"]
	if found {
		return deleted.(string)
	}
	return ""
}

func (req *PicturesCountRequest) SetPicId(picId uint64) {
	req.Request.Params["picture_id"] = picId
}

func (req *PicturesCountRequest) GetPicId() uint64 {
	picId, found := req.Request.Params["picture_id"]
	if found {
		return picId.(uint64)
	}
	return 0
}

func (req *PicturesCountRequest) SetClientType(clientType string) {
	req.Request.Params["client_type"] = clientType
}

func (req *PicturesCountRequest) GetClientType() string {
	clientType, found := req.Request.Params["client_type"]
	if found {
		return clientType.(string)
	}
	return ""
}
