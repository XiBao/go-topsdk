package dmp

type TagRequirement struct {
	Id         uint64   `json:"id,omitempty" codec:",omitempty"`         // 需求id
	Content    string   `json:"content,omitempty" codec:",omitempty"`    // 内容
	Title      string   `json:"title,omitempty" codec:",omitempty"`      // 标题
	Type       uint     `json:"type,omitempty" codec:",omitempty"`       // 类型
	Message    string   `json:"message,omitempty" codec:",omitempty"`    //错误消息
	CreateTime string   `json:"createtime,omitempty" codec:",omitempty"` //创建时间
	UpdateTime string   `json:"updatetime,omitempty" codec:",omitempty"` //创建时间
	Status     uint8    `json:"status,omitempty" codec:",omitempty"`     // 需求状态
	TagIds     []uint64 `json:"tag_ids,omitempty" codec:",omitempty"`    // 标签id
	Attributes string   `json:"attributes,omitempty" codec:",omitempty"` // 扩展属性
}
