package model

import "time"

// 合同领域模型
type Contract struct {
	ID         int64     // 主键
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间

	Name       string         // 名称
	Content    string         // 内容
	ColumnID   int64          // 专栏id
	AuthorID   int64          // 作者id
	Status     ContractStatus // 状态 0/1/2 未生效/已生效/已禁用
	Percentage float64        // 分成比例
	PaymentDay string         // 账期日
}
