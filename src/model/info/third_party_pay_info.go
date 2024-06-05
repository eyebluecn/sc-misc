package info

// 第三方支付信息
type ThirdPartyPayInfo struct {
	OrderNo string //订单号

	ThirdTransactionNo string //支付的一些token及信息

	NonceStr string //支付时候的验证信息等
}
