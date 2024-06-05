package info

import "github.com/eyebluecn/sc-misc/src/model/do"

// 支付单准备信息
type PreparePaymentInfo struct {
	Payment            *do.PaymentDO // 支付单
	ThirdTransactionNo string        // 支付的一些token及信息
	NonceStr           string        // 支付时候的验证信息等
}
