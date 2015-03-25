package opendsp

// 推广单元定向
type AdgroupTargeting struct {
	Id             uint64 `json:"id,omitempty" codec:",omitempty"`              // 定向主键id
	DspId          uint64 `json:"dsp_id,omitempty" codec:",omitempty"`          // dsp id
	DspCustId      string `json:"dsp_cust_id,omitempty" codec:",omitempty"`     // dsp系统中的推广主id
	CampaignId     uint64 `json:"campaign_id,omitempty" codec:",omitempty"`     // 推广计划id
	AdgroupId      uint64 `json:"adgroup_id,omitempty" codec:",omitempty"`      // 推广单元id
	TargetingType  uint8  `json:"targeting_type,omitempty" codec:",omitempty"`  // 定向类型
	TargetingValue string `json:"targeting_value,omitempty" codec:",omitempty"` // 定向值，多个值以英文逗号分隔","
	BidType        uint8  `json:"bid_type,omitempty" codec:",omitempty"`        // 价格出价类型：1.基础出价 2.溢价
	Price          uint   `json:"price,omitempty" codec:",omitempty"`           // 定向出价或者溢价价格，单位是分
	CreateTime     string `json:"gmt_create,omitempty" codec:",omitempty"`      // 创建时间
	ModifiedTime   string `json:"gmt_modified,omitempty" codec:",omitempty"`    // 修改时间
}
