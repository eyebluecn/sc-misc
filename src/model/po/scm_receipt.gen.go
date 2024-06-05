// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package po

import (
	"time"
)

const TableNameReceiptPO = "scm_receipt"

// ReceiptPO 收款单，作者从平台收款的单据
type ReceiptPO struct {
	ID                 int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                          // 主键
	CreateTime         time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime         time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
	OrderNo            string    `gorm:"column:order_no;not null;comment:订单编号" json:"order_no"`                                 // 订单编号
	Method             string    `gorm:"column:method;not null;comment:收款方式 ALIPAY/WEIXIN 支付宝/微信" json:"method"`                // 收款方式 ALIPAY/WEIXIN 支付宝/微信
	ThirdTransactionNo string    `gorm:"column:third_transaction_no;not null;comment:收款平台订单号" json:"third_transaction_no"`      // 收款平台订单号
	Amount             int64     `gorm:"column:amount;not null;comment:金额(单位：分)" json:"amount"`                                 // 金额(单位：分)
	Status             int32     `gorm:"column:status;not null;comment:支付状态 0/1/2 未收款/已收款/已关闭" json:"status"`                   // 支付状态 0/1/2 未收款/已收款/已关闭
}

// TableName ReceiptPO's table name
func (*ReceiptPO) TableName() string {
	return TableNameReceiptPO
}