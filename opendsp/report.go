package opendsp

type CampaignReport struct {
	DspId      uint64 `json:"dsp_id,omitempty" codec:",omitempty"`      // dspid
	DspCustId  string `json:"dsp_cust_id,omitempty" codec:",omitempty"` // dsp系统中的推广主id
	CampaignId uint64 `json:"campaign_id,omitempty" codec:",omitempty"` // 推广计划id
	RecordOn   string `json:"thedate,omitempty" codec:",omitempty"`     // 日期
	Charge     uint   `json:"charge,omitempty" codec:",omitempty"`      // 消耗
	Click      uint   `json:"click,omitempty" codec:",omitempty"`       // 点击
	Pv         uint64 `json:"pv,omitempty" codec:",omitempty"`          // 展现量
}

type AdgroupReport struct {
	DspId     uint64 `json:"dsp_id,omitempty" codec:",omitempty"`      // dspid
	DspCustId string `json:"dsp_cust_id,omitempty" codec:",omitempty"` // dsp系统中的推广主id
	AdgroupId uint64 `json:"adgroup_id,omitempty" codec:",omitempty"`  // 推广单元id
	RecordOn  string `json:"thedate,omitempty" codec:",omitempty"`     // 日期
	Charge    uint   `json:"charge,omitempty" codec:",omitempty"`      // 消耗
	Click     uint   `json:"click,omitempty" codec:",omitempty"`       // 点击
	Pv        uint64 `json:"pv,omitempty" codec:",omitempty"`          // 展现量
}

type CreativeReport struct {
	DspId      uint64 `json:"dsp_id,omitempty" codec:",omitempty"`      // dspid
	DspCustId  string `json:"dsp_cust_id,omitempty" codec:",omitempty"` // dsp系统中的推广主id
	CreativeId uint64 `json:"creative_id,omitempty" codec:",omitempty"` // 创意id
	RecordOn   string `json:"thedate,omitempty" codec:",omitempty"`     // 日期
	Charge     uint   `json:"charge,omitempty" codec:",omitempty"`      // 消耗
	Click      uint   `json:"click,omitempty" codec:",omitempty"`       // 点击
	Pv         uint64 `json:"pv,omitempty" codec:",omitempty"`          // 展现量
}

type CrowdReport struct {
	DspId     uint64 `json:"dsp_id,omitempty" codec:",omitempty"`      // dspid
	DspCustId string `json:"dsp_cust_id,omitempty" codec:",omitempty"` // dsp系统中的推广主id
	CrowdId   uint64 `json:"crowd_id,omitempty" codec:",omitempty"`    // 人群id
	RecordOn  string `json:"thedate,omitempty" codec:",omitempty"`     // 日期
	Charge    uint   `json:"charge,omitempty" codec:",omitempty"`      // 消耗
	Click     uint   `json:"click,omitempty" codec:",omitempty"`       // 点击
	Pv        uint64 `json:"pv,omitempty" codec:",omitempty"`          // 展现量
}
