package dmp

type TagOption struct {
	Id       uint64 // ID
	Name     string // 标签选项显示名称
	Value    string // 标签选项取值,用于创建人群等
	ParentId uint64 // 父选项ID,0表示是顶层选项
	SortNum  uint   // 选项排序
}
