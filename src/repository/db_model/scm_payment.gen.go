// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db_model

import (
	"time"
)

const TableNamePaymentDO = "scm_payment"

// PaymentDO 支付单，用户支付给平台的单据
type PaymentDO struct {
	ID                 int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                          // 主键
	CreateTime         time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime         time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
	OrderNo            string    `gorm:"column:order_no;not null;comment:订单编号" json:"order_no"`                                 // 订单编号
	Method             string    `gorm:"column:method;not null;comment:支付方式 ALIPAY/WEIXIN 支付宝/微信" json:"method"`                // 支付方式 ALIPAY/WEIXIN 支付宝/微信
	ThirdTransactionNo string    `gorm:"column:third_transaction_no;comment:支付平台订单号" json:"third_transaction_no"`               // 支付平台订单号
	Amount             int64     `gorm:"column:amount;not null;comment:金额(单位：分)" json:"amount"`                                 // 金额(单位：分)
	Status             int32     `gorm:"column:status;not null;comment:支付状态 0/1/2  未支付/已支付/已关闭" json:"status"`                  // 支付状态 0/1/2  未支付/已支付/已关闭
}

// TableName PaymentDO's table name
func (*PaymentDO) TableName() string {
	return TableNamePaymentDO
}
