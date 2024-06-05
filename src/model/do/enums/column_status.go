package enums

type ColumnStatus int32

const (
	ColumnStatusNew      ColumnStatus = 0 //未发布
	ColumnStatusOk       ColumnStatus = 1 //已生效
	ColumnStatusDisabled ColumnStatus = 2 //已禁用
)
