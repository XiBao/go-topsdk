package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type SimbaCreativeUpdateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaCreativeUpdateRequest() (req *SimbaCreativeUpdateRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.creative.update", Params: make(map[string]interface{}, 4)}
	req = &SimbaCreativeUpdateRequest{
		Request: &request,
	}
	return
}

//主人昵称
func (req *SimbaCreativeUpdateRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *SimbaCreativeUpdateRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}

func (req *SimbaCreativeUpdateRequest) SetAdgroupId(adgroupId uint64) {
	req.Request.Params["adgroup_id"] = adgroupId
}

func (req *SimbaCreativeUpdateRequest) GetAdgroupId() uint64 {
	adgroupId, found := req.Request.Params["adgroup_id"]
	if found {
		return adgroupId.(uint64)
	}
	return 0
}

func (req *SimbaCreativeUpdateRequest) SetCreativeId(creativeId uint64) {
	req.Request.Params["creative_id"] = creativeId
}

func (req *SimbaCreativeUpdateRequest) GetCreativeId() uint64 {
	creativeId, found := req.Request.Params["creative_id"]
	if found {
		return creativeId.(uint64)
	}
	return 0
}

func (req *SimbaCreativeUpdateRequest) SetImgUrl(imgUrl string) {
	req.Request.Params["img_url"] = imgUrl
}

func (req *SimbaCreativeUpdateRequest) GetImgUrl() string {
	imgUrl, found := req.Request.Params["img_url"]
	if found {
		return imgUrl.(string)
	}
	return ""
}

func (req *SimbaCreativeUpdateRequest) SetTitle(title string) {
	req.Request.Params["title"] = title
}

func (req *SimbaCreativeUpdateRequest) GetTitle() string {
	title, found := req.Request.Params["title"]
	if found {
		return title.(string)
	}
	return ""
}

func (req *SimbaCreativeUpdateRequest) SetPictureId(pictureId uint64) {
	req.Request.Params["picture_id"] = pictureId
}

func (req *SimbaCreativeUpdateRequest) GetPictureId() uint64 {
	pictureId, found := req.Request.Params["picture_id"]
	if found {
		return pictureId.(uint64)
	}
	return 0
}
