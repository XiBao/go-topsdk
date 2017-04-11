package TopRequest

import (
	"github.com/XiBao/topsdk"
	"time"
)

type PicturesGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewPicturesGetRequest() (req *PicturesGetRequest) {
	request := topsdk.Request{MethodName: "taobao.picture.pictures.get", Params: make(map[string]interface{})}
	req = &PicturesGetRequest{
		Request: &request,
	}
	return
}

func (req *PicturesGetRequest) SetStartDate(startDate time.Time) {
	req.Request.Params["start_date"] = startDate
}

func (req *PicturesGetRequest) GetStartDate() time.Time {
	startDate, found := req.Request.Params["start_date"]
	if found {
		return startDate.(time.Time)
	}
	return time.Time{}
}

//图片分类ID
func (req *PicturesGetRequest) SetPicCateId(picCateId uint64) {
	req.Request.Params["picture_category_id"] = picCateId
}

func (req *PicturesGetRequest) GetPicCateId() uint64 {
	picCateId, found := req.Request.Params["picture_category_id"]
	if found {
		return picCateId.(uint64)
	}
	return 0
}

func (req *PicturesGetRequest) SetOrderBy(orderBy string) {
	req.Request.Params["order_by"] = orderBy
}

func (req *PicturesGetRequest) GetOrderBy() string {
	orderBy, found := req.Request.Params["order_by"]
	if found {
		return orderBy.(string)
	}
	return ""
}

func (req *PicturesGetRequest) SetTitle(title string) {
	req.Request.Params["title"] = title
}

func (req *PicturesGetRequest) GetTitle() string {
	title, found := req.Request.Params["title"]
	if found {
		return title.(string)
	}
	return ""
}

func (req *PicturesGetRequest) SetPageSize(pageSize uint) {
	req.Request.Params["page_size"] = pageSize
}

func (req *PicturesGetRequest) GetPageSize() uint {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(uint)
	}
	return 40
}

func (req *PicturesGetRequest) SetEndDate(endDate time.Time) {
	req.Request.Params["end_date"] = endDate
}

func (req *PicturesGetRequest) GetEndDate() time.Time {
	endDate, found := req.Request.Params["end_date"]
	if found {
		return endDate.(time.Time)
	}
	return time.Time{}
}

func (req *PicturesGetRequest) SetCurrentPage(currentPage uint) {
	req.Request.Params["current_page"] = currentPage
}

func (req *PicturesGetRequest) GetCurrentPage() uint {
	currentPage, found := req.Request.Params["current_page"]
	if found {
		return currentPage.(uint)
	}
	return 1
}

func (req *PicturesGetRequest) SetStartModifiedDate(startModifiedDate time.Time) {
	req.Request.Params["start_modified_date"] = startModifiedDate
}

func (req *PicturesGetRequest) GetStartModifiedDate() time.Time {
	startModifiedDate, found := req.Request.Params["start_modified_date"]
	if found {
		return startModifiedDate.(time.Time)
	}
	return time.Time{}
}

func (req *PicturesGetRequest) SetDeleted(deleted string) {
	req.Request.Params["deleted"] = deleted
}

func (req *PicturesGetRequest) GetDeleted() string {
	deleted, found := req.Request.Params["deleted"]
	if found {
		return deleted.(string)
	}
	return ""
}

func (req *PicturesGetRequest) SetPicId(picId uint64) {
	req.Request.Params["picture_id"] = picId
}

func (req *PicturesGetRequest) GetPicId() uint64 {
	picId, found := req.Request.Params["picture_id"]
	if found {
		return picId.(uint64)
	}
	return 0
}

func (req *PicturesGetRequest) SetClientType(clientType string) {
	req.Request.Params["client_type"] = clientType
}

func (req *PicturesGetRequest) GetClientType() string {
	clientType, found := req.Request.Params["client_type"]
	if found {
		return clientType.(string)
	}
	return ""
}

func (req *PicturesGetRequest) SetUrls(urls string) {
	req.Request.Params["urls"] = urls
}

func (req *PicturesGetRequest) GetUrls() string {
	urls, found := req.Request.Params["urls"]
	if found {
		return urls.(string)
	}
	return ""
}

func (req *PicturesGetRequest) SetIsHttps(isHttps bool) {
	req.Request.Params["is_https"] = isHttps
}

func (req *PicturesGetRequest) GetIsHttps() bool {
	isHttps, found := req.Request.Params["is_https"]
	if found {
		return isHttps.(bool)
	}
	return false
}

