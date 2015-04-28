package opendsp

// 查询参数
type CampaignReportParams struct {
	StartTime      string   `json:"start_time"`                                    // 查询开始时间
	EndTime        string   `json:"end_time"`                                      // 查询结束时间
	CampaignIdList []uint64 `json:"campaign_id_list,omitempty" codec:",omitempty"` // 推广计划id列表
	Page           uint     `json:"page,omitempty" codec:",omitempty"`             // 页码
	PageSize       uint     `json:"page_size,omitempty" codec:",omitempty"`        // 每页大小
	DspCustId      string   `json:"dsp_cust_id,omitempty" codec:",omitempty"`      // dsp推广主id，建议用数字id
}

// 查询参数
type AdgroupReportParams struct {
	StartTime     string   `json:"start_time"`                                   // 查询开始时间
	EndTime       string   `json:"end_time"`                                     // 查询结束时间
	AdgroupIdList []uint64 `json:"adgroup_id_list,omitempty" codec:",omitempty"` // 推广单元id列表
	Page          uint     `json:"page,omitempty" codec:",omitempty"`            // 页码
	PageSize      uint     `json:"page_size,omitempty" codec:",omitempty"`       // 每页大小
	DspCustId     string   `json:"dsp_cust_id,omitempty" codec:",omitempty"`     // dsp推广主id，建议用数字id
}

// 查询参数
type CreativeReportParams struct {
	StartTime      string   `json:"start_time"`                                    // 查询开始时间
	EndTime        string   `json:"end_time"`                                      // 查询结束时间
	CreativeIdList []uint64 `json:"creative_id_list,omitempty" codec:",omitempty"` // 创意id列表
	Page           uint     `json:"page,omitempty" codec:",omitempty"`             // 页码
	PageSize       uint     `json:"page_size,omitempty" codec:",omitempty"`        // 每页大小
	DspCustId      string   `json:"dsp_cust_id,omitempty" codec:",omitempty"`      // dsp推广主id，建议用数字id
}

// 查询参数
type CrowdReportParams struct {
	StartTime   string   `json:"start_time"`                                 // 查询开始时间
	EndTime     string   `json:"end_time"`                                   // 查询结束时间
	CrowdIdList []uint64 `json:"crowd_id_list,omitempty" codec:",omitempty"` // 人群id列表
	Page        uint     `json:"page,omitempty" codec:",omitempty"`          // 页码
	PageSize    uint     `json:"page_size,omitempty" codec:",omitempty"`     // 每页大小
	DspCustId   string   `json:"dsp_cust_id,omitempty" codec:",omitempty"`   // dsp推广主id，建议用数字id
}
