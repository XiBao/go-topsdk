package opendsp

type Dict struct {
	KeyId      string `json:"key_id,omitempty" codec:",omitempty"`      // key_id
	KeyCode    string `json:"key_code,omitempty" codec:",omitempty"`    // key_code
	KeyName    string `json:"key_name,omitempty" codec:",omitempty"`    // key名称
	KeyLevel   uint   `json:"key_level,omitempty" codec:",omitempty"`   // 层级
	ParentCode string `json:"parent_code,omitempty" codec:",omitempty"` // 上层的节点code，如果第一级则为空
	KeyDesc    string `json:"key_desc,omitempty" codec:",omitempty"`    // 描述
}
