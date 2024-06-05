// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package po

import (
	"time"
)

const TableNameContractPO = "scm_contract"

// ContractPO 作者表
type ContractPO struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                          // 主键
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
	Name       string    `gorm:"column:name;not null;comment:名称" json:"name"`                                           // 名称
	Content    string    `gorm:"column:content;not null;comment:内容" json:"content"`                                     // 内容
	ColumnID   int64     `gorm:"column:column_id;not null;comment:专栏id" json:"column_id"`                               // 专栏id
	AuthorID   int64     `gorm:"column:author_id;not null;comment:作者id" json:"author_id"`                               // 作者id
	Status     int32     `gorm:"column:status;not null;comment:状态 0/1/2 未生效/已生效/已禁用" json:"status"`                     // 状态 0/1/2 未生效/已生效/已禁用
	Percentage float64   `gorm:"column:percentage;comment:分成比例" json:"percentage"`                                      // 分成比例
	PaymentDay string    `gorm:"column:payment_day;comment:账期日" json:"payment_day"`                                     // 账期日
}

// TableName ContractPO's table name
func (*ContractPO) TableName() string {
	return TableNameContractPO
}