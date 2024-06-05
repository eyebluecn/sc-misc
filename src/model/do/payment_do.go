package do

import (
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"time"
)

// 支付单领域模型
type PaymentDO struct {
	ID                 int64               // 主键
	CreateTime         time.Time           // 创建时间
	UpdateTime         time.Time           // 更新时间
	OrderNo            string              // 订单编号
	Method             string              // 支付方式 ALIPAY/WEIXIN 支付宝/微信
	ThirdTransactionNo string              // 支付平台订单号
	Amount             int64               // 金额(单位：分)
	Status             enums.PaymentStatus // 支付状态 0/1/2  未支付/已支付/已关闭
}
