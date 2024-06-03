package model

import "time"

// 小编领域模型
type Editor struct {
	ID         int64     // 主键
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间

	Username string // 昵称
	Password string // 密码
	WorkNo   string // 工号

}
