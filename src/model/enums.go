package model

type ColumnStatus int32

const (
	ColumnStatusNew      ColumnStatus = 0 //未发布
	ColumnStatusOk       ColumnStatus = 1 //已生效
	ColumnStatusDisabled ColumnStatus = 2 //已禁用
)

type ColumnQuoteStatus int32

const (
	ColumnQuoteStatusNew ColumnQuoteStatus = 0 //未生效
	ColumnQuoteStatusOk  ColumnQuoteStatus = 1 //已生效
)

type ContractStatus int32

const (
	ContractStatusNew      ContractStatus = 0 //未发布
	ContractStatusOk       ContractStatus = 1 //已生效
	ContractStatusDisabled ContractStatus = 2 //已禁用
)

type PaymentStatus int32

const (
	PaymentStatusUnpaid PaymentStatus = 0 //未支付
	PaymentStatusPaid   PaymentStatus = 1 //已支付
	PaymentStatusClosed PaymentStatus = 2 //已关闭
)
