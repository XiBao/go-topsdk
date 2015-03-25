package dmp

type DmpCrowd struct {
	CrowdId           uint64                `json:"crowd_id,omitempty" codec:",omitempty"`           // 人群ID
	UserId            uint64                `json:"user_id,omitempty" codec:",omitempty"`            // 用户ID
	CrowdName         string                `json:"crowd_name,omitempty" codec:",omitempty"`         // 人群名称
	AppKey            uint64                `json:"app_key,omitempty" codec:",omitempty"`            // app_key
	selects           []*DmpSelectTagOption `json:"selects,omitempty" codec:",omitempty"`            // 选项
	Coverage          uint                  `json:"coverage,omitempty" codec:",omitempty"`           // 覆盖人数
	MemberId          uint64                `json:"member_id,omitempty" codec:",omitempty"`          // 所属用户ID
	LookalikeMultiple uint                  `json:"lookalike_multiple,omitempty" codec:",omitempty"` // 放大倍数
}
