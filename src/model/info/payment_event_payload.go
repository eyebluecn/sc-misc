package info

import (
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"time"
)

// 领域事件的payload信息
type PaymentEventPayload struct {
	// 支付单 ID
	PaymentId int64
	// 订单编号
	OrderNo string
	// 支付方式（如"ALIPAY"表示支付宝，"WEIXIN"表示微信）
	Method string
	// 金额（单位：分）
	Amount int64
	// 支付状态（如"UNPAID"表示未支付，"PAID"表示已支付，"CLOSED"表示已关闭）
	Status enums.PaymentStatus
	// 发生时间
	OccurTime time.Time
}
