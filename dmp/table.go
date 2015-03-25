package dmp

type Table struct {
	Id                uint64    `json:"id,omitempty" codec:",omitempty"`                  // 表Id
	DpId              uint64    `json:"dp_id,omitempty" codec:",omitempty"`               // dpid
	AppId             string    `json:"app_id,omitempty" codec:",omitempty"`              // app_id
	Name              string    `json:"name,omitempty" codec:",omitempty"`                // 名字
	Comment           string    `json:"comment,omitempty" codec:",omitempty"`             // 表注释
	PartitionKey      string    `json:"partition_key,omitempty" codec:",omitempty"`       // 分区key
	PartitionValue    string    `json:"partition_value,omitempty" codec:",omitempty"`     // 分区值
	ProjectCode       string    `json:"project_code,omitempty" codec:",omitempty"`        // 项目code
	SecLevel          uint8     `json:"sec_level,omitempty" codec:",omitempty"`           // 安全等级，分为4等级
	State             uint      `json:"state,omitempty" codec:",omitempty"`               // 表状态，上传中，识别中等
	Columns           []*Column `json:"columns,omitempty" codec:",omitempty"`             // 列信息
	AllParitionValues []string  `json:"all_parition_values,omitempty" codec:",omitempty"` // 分区值
	Code              string    `json:"code,omitempty" codec:",omitempty"`                // 表英文名
	Type              uint      `json:"type,omitempty" codec:",omitempty"`                // 原始表，识别表，输出表等
	OdpsType          uint      `json:"odps_type,omitempty" codec:",omitempty"`           // odps 类型
}
