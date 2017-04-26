package TopRequest

import (
	"github.com/XiBao/topsdk"
	"strconv"
	"strings"
)

type ZuanshiBannerInterestFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerInterestFindRequest() (req *ZuanshiBannerInterestFindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.interest.find", Params: make(map[string]interface{}, 2)}
	req = &ZuanshiBannerInterestFindRequest{
		Request: &request,
	}
	return
}

// 按宝贝ID列表查询兴趣点, 最大列表长度：20
func (req *ZuanshiBannerInterestFindRequest) SetItemIds(ids []uint64) {
	var idsStr []string
	for _, id := range ids {
		idsStr = append(idsStr, strconv.FormatUint(id, 10))
	}
	req.Request.Params["item_ids"] = idsStr
}

func (req *ZuanshiBannerInterestFindRequest) GetItemIds() []uint64 {
	var ids []uint64
	if idStr, found := req.Request.Params["item_ids"]; found {
		idsStr := strings.Split(idStr.(string), ",")
		for _, id := range idsStr {
			if aId, err := strconv.ParseUint(id, 10, 64); err == nil && aId > 0 {
				ids = append(ids, aId)
			}
		}
	}
	return ids
}

// 按店铺名称查询兴趣点
func (req *ZuanshiBannerInterestFindRequest) SetNickname(nickname string) {
	req.Request.Params["nickname"] = nickname
}

func (req *ZuanshiBannerCampaignStatusRequest) GetNickname() string {
	if nickname, found := req.Request.Params["nickname"]; found {
		return nickname.(string)
	}
	return ""
}

// 按关键字查询兴趣点
func (req *ZuanshiBannerInterestFindRequest) SetKeyword(keyword string) {
	req.Request.Params["keyword"] = keyword
}

func (req *ZuanshiBannerCampaignStatusRequest) GetKeyword() string {
	if keyword, found := req.Request.Params["keyword"]; found {
		return keyword.(string)
	}
	return ""
}
