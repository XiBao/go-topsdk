package dmp

type DmpTag struct {
	Id              uint64               `json:"id,omitempty" codec:",omitempty"`                    // 标签ID
	ChartType       string               `json:"chart_type,omitempty" codec:",omitempty"`            // 图标类型
	TagType         string               `json:"tag_type,omitempty" codec:",omitempty"`              // 标签类型
	Price           float32              `json:"price,omitempty" codec:",omitempty"`                 // 价格
	TagDesc         string               `json:"tag_desc,omitempty" codec:",omitempty"`              // 标签描述
	ValidDate       string               `json:"valid_date,omitempty" codec:",omitempty"`            // 过期时间
	TagOptions      []*DmpTagOption      `json:"tag_option_dtos,omitempty" codec:",omitempty"`       // 标签选项集合
	TagName         string               `json:"tag_name,omitempty" codec:",omitempty"`              // 标签名称
	TagOptionGroups []*DmpTagOptionGroup `json:"tag_option_group_dtos,omitempty" codec:",omitempty"` // 标签选项组集合
}
