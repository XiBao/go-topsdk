package opendsp

// 推广单元
type Adgroup struct {
	Id             uint64 `json:"id,omitempty" codec:",omitempty"`              // adgroup的主键id
	DspCustId      string `json:"dsp_cust_id"`                                  // dsp系统中的推广主id，是string类型
	CampaignId     uint64 `json:"campaign_id"`                                  // campaign计划id
	Title          string `json:"title"`                                        // adgroup的名称
	DspId          uint64 `json:"dsp_id,omitempty" codec:",omitempty"`          // dsp id
	DspAdgroupId   uint64 `json:"dsp_adgroup_id,omitempty" codec:",omitempty"`  // dsp系统中的推广单元id
	Status         uint8  `json:"status,omitempty" codec:",omitempty"`          // 上下线状态：0表示下线 1表示上线
	CreateTime     string `json:"gmt_create,omitempty" codec:",omitempty"`      // 创建时间
	ModifiedTime   string `json:"gmt_modified,omitempty" codec:",omitempty"`    // 修改时间
	FreqImpression int8   `json:"freq_impression,omitempty" codec:",omitempty"` // 用户频次控制（-1表示没有限制）
}
