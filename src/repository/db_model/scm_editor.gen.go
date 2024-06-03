// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db_model

import (
	"time"
)

const TableNameEditorDO = "scm_editor"

// EditorDO 小二表
type EditorDO struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                          // 主键
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
	Username   string    `gorm:"column:username;not null;comment:昵称" json:"username"`                                   // 昵称
	Password   string    `gorm:"column:password;not null;comment:密码" json:"password"`                                   // 密码
	WorkNo     string    `gorm:"column:work_no;not null;comment:工号" json:"work_no"`                                     // 工号
}

// TableName EditorDO's table name
func (*EditorDO) TableName() string {
	return TableNameEditorDO
}
