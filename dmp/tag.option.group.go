package dmp

type TagOptionGroup struct {
	Id      uint64       `json:",omitempty" codec:",omitempty"` // ID
	Name    string       `json:",omitempty" codec:",omitempty"` // 名称
	Type    string       `json:",omitempty" codec:",omitempty"` // 标签类型，单选，多选，输入框等
	Options []*TagOption `json:",omitempty" codec:",omitempty"` // tagOptionDTOs
}
