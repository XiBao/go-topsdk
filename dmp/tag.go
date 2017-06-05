package dmp

import (
	"time"
)

type Tag struct {
	Id           uint64            `json:",omitempty" codec:",omitempty"` // 标签ID
	Desc         string            `json:",omitempty" codec:",omitempty"` // 标签描述
	Name         string            `json:",omitempty" codec:",omitempty"` // 标签名称
	OptionGroups []*TagOptionGroup `json:",omitempty" codec:",omitempty"` // 标签选项组集合
	Share        uint8             `json:",omitempty" codec:",omitempty"` // 1-共享标签,0-私有标签
	Status       uint8             `json:",omitempty" codec:",omitempty"` // 1-正常状态,其它都为不可用状态
	ValidDate    time.Time         `json:",omitempty" codec:",omitempty"` // 过期时间
}
