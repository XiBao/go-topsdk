package dmp

type DmpSelectTagOption struct {
	Values        string `json:"values,omitempty" codec:",omitempty"`          // 值
	OptionGroupId uint64 `json:"option_group_id,omitempty" codec:",omitempty"` // 值分组,在标签列表返中获取
	TagId         uint64 `json:"tag_id,omitempty" codec:",omitempty"`          // 标签ID
}
