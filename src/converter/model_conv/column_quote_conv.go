package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
)

// 转为枚举
func ConvertColumnQuoteStatus(status int32) model.ColumnQuoteStatus {
	return model.ColumnQuoteStatus(status)
}

// 数据库模型转换为领域模型
func ConvertColumnQuote(thing *db_model.ColumnQuoteDO) *model.ColumnQuote {
	if thing == nil {
		return nil
	}

	return &model.ColumnQuote{
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
func ConvertColumnQuotes(things []*db_model.ColumnQuoteDO) []*model.ColumnQuote {
	if things == nil {
		return nil
	}
	var readers []*model.ColumnQuote
	for _, item := range things {
		readers = append(readers, ConvertColumnQuote(item))
	}
	return readers
}
