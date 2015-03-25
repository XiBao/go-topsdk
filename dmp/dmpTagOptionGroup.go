package dmp

type DmpTagOptionGroup struct {
	Id        uint64 `json:"id,omitempty" codec:",omitempty"`         // ID
	GroupName string `json:"group_name,omitempty" codec:",omitempty"` // 名称
	Type      string `json:"type,omitempty" codec:",omitempty"`       // 类型
}
