package enums

type ColumnQuoteStatus int32

const (
	ColumnQuoteStatusNew ColumnQuoteStatus = 0 //未生效
	ColumnQuoteStatusOk  ColumnQuoteStatus = 1 //已生效
)
