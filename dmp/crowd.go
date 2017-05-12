package dmp

import (
	"time"
)

type Crowd struct {
	ValidDate        time.Time // 人群过期时间
	EnableTime       time.Time // 人群生效时间
	UpdateTime       time.Time // 创建时间
	CreateTime       time.Time //更新时间
	SelectTagOptions []*SelectTagOption
	Id               uint64 // 人群ID
	Name             string // 人群名称
	FullStatus       uint   // 2-过期失效, 11-同步中, 12-已同步
	Coverage         uint64 // 覆盖人数
}
