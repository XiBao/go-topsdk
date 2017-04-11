package TopRequest

import (
	"github.com/XiBao/topsdk"
	"time"
)

type PictureGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewPictureGetRequest() (req *PictureGetRequest) {
	request := topsdk.Request{MethodName: "taobao.picture.pictures.get", Params: make(map[string]interface{})}
	req = &PictureGetRequest{
		Request: &request,
	}
	return
}

func (req *PictureGetRequest) SetPicId(picId uint64) {
	req.Request.Params["picture_id"] = picId
}

func (req *PictureGetRequest) GetPicId() uint64 {
	picId, found := req.Request.Params["picture_id"]
	if found {
		return picId.(uint64)
	}
	return 0
}

//图片分类ID
func (req *PictureGetRequest) SetPicCateId(picCateId uint64) {
	req.Request.Params["picture_category_id"] = picCateId
}

func (req *PictureGetRequest) GetPicCateId() uint64 {
	picCateId, found := req.Request.Params["picture_category_id"]
	if found {
		return picCateId.(uint64)
	}
	return 0
}

func (req *PictureGetRequest) SetDeleted(deleted string) {
	req.Request.Params["deleted"] = deleted
}

func (req *PictureGetRequest) GetDeleted() string {
	deleted, found := req.Request.Params["deleted"]
	if found {
		return deleted.(string)
	}
	return ""
}

func (req *PictureGetRequest) SetTitle(title string) {
	req.Request.Params["title"] = title
}

func (req *PictureGetRequest) GetTitle() string {
	title, found := req.Request.Params["title"]
	if found {
		return title.(string)
	}
	return ""
}

func (req *PictureGetRequest) SetOrderBy(orderBy string) {
	req.Request.Params["order_by"] = orderBy
}

func (req *PictureGetRequest) GetOrderBy() string {
	orderBy, found := req.Request.Params["order_by"]
	if found {
		return orderBy.(string)
	}
	return ""
}

func (req *PictureGetRequest) SetStartDate(startDate time.Time) {
	req.Request.Params["start_date"] = startDate
}

func (req *PictureGetRequest) GetStartDate() time.Time {
	startDate, found := req.Request.Params["start_date"]
	if found {
		return startDate.(time.Time)
	}
	return time.Time{}
}

func (req *PictureGetRequest) SetEndDate(endDate time.Time) {
	req.Request.Params["end_date"] = endDate
}

func (req *PictureGetRequest) GetEndDate() time.Time {
	endDate, found := req.Request.Params["end_date"]
	if found {
		return endDate.(time.Time)
	}
	return time.Time{}
}

func (req *PictureGetRequest) SetPageNo(pageNo uint) {
	req.Request.Params["current_page"] = pageNo
}

func (req *PictureGetRequest) GetPageNo() uint {
	pageNo, found := req.Request.Params["page_no"]
	if found {
		return pageNo.(uint)
	}
	return 1
}

func (req *PictureGetRequest) SetPageSize(pageSize uint) {
	req.Request.Params["page_size"] = pageSize
}

func (req *PictureGetRequest) GetPageSize() uint {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(uint)
	}
	return 40
}

func (req *PictureGetRequest) SetModifiedTime(modifiedTime time.Time) {
	req.Request.Params["modified_time"] = modifiedTime
}

func (req *PictureGetRequest) GetModifiedTime() time.Time {
	modifiedTime, found := req.Request.Params["modified_time"]
	if found {
		return modifiedTime.(time.Time)
	}
	return time.Time{}
}

func (req *PictureGetRequest) SetClientType(clientType string) {
	req.Request.Params["client_type"] = clientType
}

func (req *PictureGetRequest) GetClientType() string {
	clientType, found := req.Request.Params["client_type"]
	if found {
		return clientType.(string)
	}
	return ""
}

func (req *PictureGetRequest) SetIsHttps(isHttps bool) {
	req.Request.Params["is_https"] = isHttps
}

func (req *PictureGetRequest) GetIsHttps() bool {
	isHttps, found := req.Request.Params["is_https"]
	if found {
		return isHttps.(bool)
	}
	return false
}

func (req *PictureGetRequest) SetUrls(urls string) {
	req.Request.Params["urls"] = urls
}

func (req *PictureGetRequest) GetUrls() string {
	urls, found := req.Request.Params["urls"]
	if found {
		return urls.(string)
	}
	return ""
}
