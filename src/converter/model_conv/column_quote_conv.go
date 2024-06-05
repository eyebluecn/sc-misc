package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 转为枚举
func ConvertColumnQuoteStatus(status int32) enums.ColumnQuoteStatus {
	return enums.ColumnQuoteStatus(status)
}

// 数据库模型转换为领域模型
func ConvertColumnQuote(thing *po.ColumnQuotePO) *do.ColumnQuote {
	if thing == nil {
		return nil
	}

	return &do.ColumnQuote{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		ColumnID:   thing.ColumnID,
		EditorID:   thing.EditorID,
		Price:      thing.Price,
		Status:     ConvertColumnQuoteStatus(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertColumnQuotes(things []*po.ColumnQuotePO) []*do.ColumnQuote {
	if things == nil {
		return nil
	}
	var readers []*do.ColumnQuote
	for _, item := range things {
		readers = append(readers, ConvertColumnQuote(item))
	}
	return readers
}
