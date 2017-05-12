package dmp

type SelectTagOptionSet struct {
	Options []*SelectTagOption `json:"selects,omitempty" codec:",omitempty"`
}

type SelectTagOption struct {
	Values        string `json:"values,omitempty" codec:",omitempty"`          // 值
	OptionGroupId uint64 `json:"option_group_id,omitempty" codec:",omitempty"` // 值分组,在标签列表返中获取
	TagId         uint64 `json:"tag_id,omitempty" codec:",omitempty"`          // 标签ID
	TagName       string `json:"tag_name,omitempty" codec:",omitempty"`        // 标签名称
	Names         string `json:"names,omitempty" codec:",omitempty"`           // 标签选项名称
}
