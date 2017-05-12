package dmp

type TagOptionGroup struct {
	Id      uint64       // ID
	Name    string       // 名称
	Type    string       // 标签类型，单选，多选，输入框等
	Options []*TagOption // tagOptionDTOs
}
