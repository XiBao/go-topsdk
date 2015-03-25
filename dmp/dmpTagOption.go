package dmp

type DmpTagOption struct {
	Id             uint64 `json:"id,omitempty" codec:",omitempty"`               // ID
	OptionName     string `json:"option_name,omitempty" codec:",omitempty"`      // 值名称
	OptionGroupId  uint64 `json:"option_group_id,omitempty" codec:",omitempty"`  // 标签分组ID
	ParentOptionId uint64 `json:"parent_option_id,omitempty" codec:",omitempty"` // 父节点ID
	SortNum        int8   `json:"sort_num,omitempty" codec:",omitempty"`         // 排序
	OptionValue    string `json:"option_value,omitempty" codec:",omitempty"`     // 值
}
