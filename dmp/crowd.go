package dmp

import (
	"time"
)

type Crowd struct {
	ValidDate        time.Time          `json:",omitempty" codec:",omitempty"` // 人群过期时间
	EnableTime       time.Time          `json:",omitempty" codec:",omitempty"` // 人群生效时间
	UpdateTime       time.Time          `json:",omitempty" codec:",omitempty"` // 创建时间
	CreateTime       time.Time          `json:",omitempty" codec:",omitempty"` //更新时间
	SelectTagOptions []*SelectTagOption `json:",omitempty" codec:",omitempty"`
	Id               uint64             `json:",omitempty" codec:",omitempty"` // 人群ID
	Name             string             `json:",omitempty" codec:",omitempty"` // 人群名称
	FullStatus       uint               `json:",omitempty" codec:",omitempty"` // 2-过期失效, 11-同步中, 12-已同步
	Coverage         uint64             `json:",omitempty" codec:",omitempty"` // 覆盖人数
}
