package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserTargetAdzoneRtrptsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserTargetAdzoneRtrptsGetRequest() (req *ZuanshiAdvertiserTargetAdzoneRtrptsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.target.adzone.rtrpts.get", Params: make(map[string]interface{}, 5)}
	req = &ZuanshiAdvertiserTargetAdzoneRtrptsGetRequest{
		Request: &request,
	}
	return
}

//查询日期
func (req *ZuanshiAdvertiserTargetAdzoneRtrptsGetRequest) SetLogDate(logDate string) {
	req.Request.Params["log_date"] = logDate
}

func (req *ZuanshiAdvertiserTargetAdzoneRtrptsGetRequest) GetLogDate() string {
	logDate, found := req.Request.Params["log_date"]
	if found {
		return logDate.(string)
	}
	return ""
}

// 计划id
func (req *ZuanshiAdvertiserTargetAdzoneRtrptsGetRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiAdvertiserTargetAdzoneRtrptsGetRequest) GetCampaignId() uint64 {
	id, found := req.Request.Params["campaign_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 推广单元id
func (req *ZuanshiAdvertiserTargetAdzoneRtrptsGetRequest) SetAdgroupId(id uint64) {
	req.Request.Params["adgroup_id"] = id
}

func (req *ZuanshiAdvertiserTargetAdzoneRtrptsGetRequest) GetAdgroupId() uint64 {
	id, found := req.Request.Params["adgroup_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 定向标签id
func (req *ZuanshiAdvertiserTargetAdzoneRtrptsGetRequest) SetTargetId(id uint64) {
	req.Request.Params["target_id"] = id
}

func (req *ZuanshiAdvertiserTargetAdzoneRtrptsGetRequest) GetTargetId() uint64 {
	id, found := req.Request.Params["target_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 资源位id
func (req *ZuanshiAdvertiserTargetAdzoneRtrptsGetRequest) SetAdzoneId(id uint64) {
	req.Request.Params["adzone_id"] = id
}

func (req *ZuanshiAdvertiserTargetAdzoneRtrptsGetRequest) GetAdzoneId() uint64 {
	id, found := req.Request.Params["adzone_id"]
	if found {
		return id.(uint64)
	}
	return 0
}
