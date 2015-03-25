package opendsp

type Report struct {
	DspId      uint64 `json:"dsp_id,omitempty" codec:",omitempty"`      // dspid
	DspCustId  string `json:"dsp_cust_id,omitempty" codec:",omitempty"` // dsp系统中的推广主id
	CampaignId uint64 `json:"campaign_id,omitempty" codec:",omitempty"` // 推广计划id
	RecordOn   string `json:"thedate,omitempty" codec:",omitempty"`     // 日期
	Charge     uint   `json:"charge,omitempty" codec:",omitempty"`      // 消耗
	Click      uint   `json:"click,omitempty" codec:",omitempty"`       // 点击
	Pv         uint   `json:"pv,omitempty" codec:",omitempty"`          // 展现量
}
