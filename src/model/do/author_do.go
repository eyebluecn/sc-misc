package do

import "time"

// 作者领域模型
type AuthorDO struct {
	ID         int64     // 主键
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间
	Username   string    // 用户名
	Password   string    // 密码
	Realname   string    // 真实姓名
}
