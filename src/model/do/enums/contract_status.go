package enums

type ContractStatus int32

const (
	ContractStatusNew      ContractStatus = 0 //未发布
	ContractStatusOk       ContractStatus = 1 //已生效
	ContractStatusDisabled ContractStatus = 2 //已禁用
)
