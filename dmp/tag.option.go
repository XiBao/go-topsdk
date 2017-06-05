package dmp

type TagOption struct {
	Id       uint64 `json:",omitempty" codec:",omitempty"` // ID
	Name     string `json:",omitempty" codec:",omitempty"` // 标签选项显示名称
	Value    string `json:",omitempty" codec:",omitempty"` // 标签选项取值,用于创建人群等
	ParentId uint64 `json:",omitempty" codec:",omitempty"` // 父选项ID,0表示是顶层选项
	SortNum  uint   `json:",omitempty" codec:",omitempty"` // 选项排序
}
