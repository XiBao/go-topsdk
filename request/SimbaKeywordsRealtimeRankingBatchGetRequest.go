package TopRequest

import (
	"github.com/XiBao/topsdk"
	"strconv"
	"strings"
)

type SimbaKeywordsRealtimeRankingBatchGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaKeywordsRealtimeRankingBatchGetRequest() (req *SimbaKeywordsRealtimeRankingBatchGetRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.keywords.realtime.ranking.batch.get", Params: make(map[string]interface{}, 3)}
	req = &SimbaKeywordsRealtimeRankingBatchGetRequest{
		Request: &request,
	}
	return
}

// 推广组id
func (req *SimbaKeywordsRealtimeRankingBatchGetRequest) SetAdGroupId(adgroup_id uint64) {
	req.Request.Params["ad_group_id"] = adgroup_id
}

func (req *SimbaKeywordsRealtimeRankingBatchGetRequest) GetAdGroupId() uint64 {
	adgroup_id, found := req.Request.Params["ad_group_id"]
	if found {
		return adgroup_id.(uint64)
	}
	return 0
}

//关键词列表集合，最多20个
func (req *SimbaKeywordsRealtimeRankingBatchGetRequest) SetBidwordIds(ids []uint64) {
	var idsStr []string
	for _, id := range ids {
		idsStr = append(idsStr, strconv.FormatUint(id, 10))
	}
	req.Request.Params["bidword_ids"] = strings.Join(idsStr, ",")
}

func (req *SimbaKeywordsRealtimeRankingBatchGetRequest) GetBidwordIds() []uint64 {
	var ids []uint64
	if idStr, found := req.Request.Params["bidword_ids"]; found {
		idsStr := strings.Split(idStr.(string), ",")
		for _, id := range idsStr {
			if aId, err := strconv.ParseUint(id, 10, 64); err == nil && aId > 0 {
				ids = append(ids, aId)
			}
		}
	}
	return ids
}

//主人昵称
func (req *SimbaKeywordsRealtimeRankingBatchGetRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *SimbaKeywordsRealtimeRankingBatchGetRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}
