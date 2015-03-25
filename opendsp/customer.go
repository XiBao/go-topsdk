package opendsp

// 推广用户
type Customer struct {
	DspCustId    string `json:"dsp_cust_id"`                               // DSP自己维护的推广主ID
	Id           uint64 `json:"id,omitempty" codec:",omitempty"`           // 主键
	DspId        uint64 `json:"dsp_id,omitempty" codec:",omitempty"`       // DSP用户ID
	Status       uint8  `json:"status,omitempty" codec:",omitempty"`       // 用户上下线状态(0:下线 1上线)
	CreateTime   string `json:"gmt_create,omitempty" codec:",omitempty"`   // 创建时间
	ModifiedTime string `json:"gmt_modified,omitempty" codec:",omitempty"` // 更新时间
}
