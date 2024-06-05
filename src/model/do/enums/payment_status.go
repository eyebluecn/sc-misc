package enums

type PaymentStatus int32

const (
	PaymentStatusUnpaid PaymentStatus = 0 //未支付
	PaymentStatusPaid   PaymentStatus = 1 //已支付
	PaymentStatusClosed PaymentStatus = 2 //已关闭
)
