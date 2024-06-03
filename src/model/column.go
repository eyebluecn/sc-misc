package model

import "time"

// 专栏领域模型
type Column struct {
	ID         int64     // 主键
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间

	Name     string       // 专栏名称
	AuthorID int64        // 作者id
	Status   ColumnStatus //状态
}
