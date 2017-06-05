package dmp

type SelectTagOptionSet struct {
	Options []*SelectTagOption `json:"selects,omitempty" codec:"selects,omitempty"`
}

type SelectTagOption struct {
	Values        string `json:"values,omitempty" codec:"values,omitempty"`                   // 值
	OptionGroupId uint64 `json:"option_group_id,omitempty" codec:"option_group_id,omitempty"` // 值分组,在标签列表返中获取
	TagId         uint64 `json:"tag_id,omitempty" codec:"tag_id,omitempty"`                   // 标签ID
	TagName       string `json:"tag_name,omitempty" codec:"tag_name,omitempty"`               // 标签名称
	Names         string `json:"names,omitempty" codec:"names,omitempty"`                     // 标签选项名称
}
