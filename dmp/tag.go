package dmp

import (
	"time"
)

type Tag struct {
	Id           uint64            // 标签ID
	Desc         string            // 标签描述
	Name         string            // 标签名称
	OptionGroups []*TagOptionGroup // 标签选项组集合
	Share        uint8             // 1-共享标签,0-私有标签
	Status       uint8             // 1-正常状态,其它都为不可用状态
	ValidDate    time.Time         // 过期时间
}
