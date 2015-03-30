package opendsp

// 推广创意
type Creative struct {
	TemplateId     uint64 `json:"template_id"`                                   // 所属模板ID
	CatId          string `json:"cat_id"`                                        // 创意所属类目ID,字典服务接口获取
	DspCustId      string `json:"dsp_cust_id"`                                   // DSP自己维护的推广用户ID
	Title          string `json:"title"`                                         // 创意标题
	Content        string `json:"content"`                                       // 创意内容
	Size           string `json:"size"`                                          // 创意尺寸
	Type           uint8  `json:"type"`                                          // 创意类型[1,2]
	Level          uint8  `json:"level"`                                         // 创意级别[1-100]
	DestinationUrl string `json:"destination_url"`                               // 声明URL
	DspCreativeId  string `json:"dsp_creative_id,omitempty", codec:",omitempty"` // DSP自己维护的创意ID
	Id             uint64 `json:"id,omitempty" codec:",omitempty"`               // 创意ID
	DspId          uint64 `json:"dsp_id,omitempty" codec:",omitempty"`           // DSP用户ID
	Status         uint8  `json:"status" codec:",omitempty"`                     // 创意上下线状态[0,1]
	CreateTime     string `json:"gmt_create,omitempty" codec:",omitempty"`       // 创建时间
	ModifiedTime   string `json:"gmt_modified,omitempty" codec:",omitempty"`     // 更新时间
}
