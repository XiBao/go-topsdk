package opendsp

// 推广计划
type Campaign struct {
	DspCustId         string `json:"dsp_cust_id"`                                     // dsp系统中的推广主id，是string类型
	Title             string `json:"title"`                                           // 计划标题
	Type              uint8  `json:"type"`                                            // 计划类型 计划类型,1.cpt 2.cpm 3.cpc
	DayBudget         uint   `json:"daybudget"`                                       // 日预算 单位是分，不限预算填-1
	StartTime         string `json:"start_time"`                                      // 计划开始时间
	EndTime           string `json:"end_time"`                                        // 计划结束时间
	DspCampaignId     uint64 `json:"dsp_campaign_id,omitempty" codec:",omitempty"`    // dsp对应的计划id
	Status            uint8  `json:"status,omitempty" codec:",omitempty"`             // 上下线状态：0表示下线 1表示上线
	OptimizationType  uint8  `json:"optimization_type,omitempty" codec:",omitempty"`  // 优化目标：1cpm 2cpc 3cpa
	OptimizationValue uint8  `json:"optimization_value,omitempty" codec:",omitempty"` // 优化值，单位是分
	Smooth            uint8  `json:"smooth,omitempty" codec:",omitempty"`             // 投放方式:1:尽快;2:平滑
	FreqImpression    int8   `json:"freq_impression,omitempty" codec:",omitempty"`    // 用户频次控制（-1表示没有限制）
	Id                uint64 `json:"id,omitempty" codec:",omitempty"`                 // 主键id
	DspId             uint64 `json:"dsp_id,omitempty" codec:",omitempty"`             // dsp id
	CreateTime        string `json:"gmt_create,omitempty" codec:",omitempty"`         // 创建时间
	ModifiedTime      string `json:"gmt_modified,omitempty" codec:",omitempty"`       // 更新时间
}
