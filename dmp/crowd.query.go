package dmp

type CrowdQuery struct {
	Name        string `json:"crowd_name,omitempty" codec:"crowd_name,omitempty"`     // 人群名称模糊查询
	QueryStatus uint   `json:"query_status,omitempty" codec:"query_status,omitempty"` // 3-未同步钻展渠道人群;2-过期失效人群,11-钻展同步中,12-钻展已同步,10-钻展同步中和已同步人群
}
