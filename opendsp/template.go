package opendsp

// 创意模板
type Template struct {
	Name             string `json:"name"`                                         // 模板名称
	ViewTemplateCode string `json:"view_template_code"`                           // 模板内容
	CreativeType     uint8  `json:"creative_type"`                                // 支持创意类型[1,2]
	Type             uint8  `json:"type"`                                         // 模板类型[1,2]
	Status           uint8  `json:"status" codec:",omitempty"`                    // 模板上下线状态[0,1]
	Size             string `json:"size"`                                         // 模板尺寸
	Id               uint64 `json:"id,omitempty" codec:",omitempty"`              // 创意模板ID
	DspId            uint64 `json:"dsp_id,omitempty" codec:",omitempty"`          // dspid
	DspTemplateId    string `json:"dsp_template_id,omitempty" codec:",omitempty"` // DSP自己维护的创意模板ID
	Description      string `json:"description,omitempty" codec:",omitempty"`     // 模板描述
	CreateTime       string `json:"gmt_create,omitempty" codec:",omitempty"`      // 创建时间
	ModifiedTime     string `json:"gmt_modified,omitempty" codec:",omitempty"`    // 更新时间
}
