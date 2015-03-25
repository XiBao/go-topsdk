package dmp

type Column struct {
	Id        uint64 `json:"id,omitempty" codec:",omitempty"`         // 列Id
	DataType  uint   `json:"data_type,omitempty" codec:",omitempty"`  // 数据类型
	SecLevel  uint8  `json:"sec_level,omitempty" codec:",omitempty"`  // 安全等级，分为4等级
	Name      string `json:"name,omitempty" codec:",omitempty"`       // 列中文名
	RecogType uint   `json:"recog_type,omitempty" codec:",omitempty"` // 识别类型，如邮箱，手机，imei, mac等
	Code      string `json:"code,omitempty" codec:",omitempty"`       // 表英文名
}
