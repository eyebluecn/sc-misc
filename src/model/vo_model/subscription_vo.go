package vo_model

import "time"

// 订阅 领域模型
type SubscriptionVO struct {
	ID         int64              // 主键
	CreateTime time.Time          // 创建时间
	UpdateTime time.Time          // 更新时间
	ReaderID   int64              // 读者id
	ColumnID   int64              // 专栏id
	OrderID    int64              // 订单id
	Status     SubscriptionStatus // 状态 0/1/2
}
