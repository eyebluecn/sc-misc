package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
)

// 枚举转为存储的整数
func ColumnQuoteStatusToStorage(status model.ColumnQuoteStatus) int32 {
	return int32(status)
}

// 数据库模型转换为领域模型
func ConvertColumnQuoteDO(thing *model.ColumnQuote) *db_model.ColumnQuoteDO {
	if thing == nil {
		return nil
	}

	return &db_model.ColumnQuoteDO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		ColumnID:   thing.ColumnID,
		EditorID:   thing.EditorID,
		Price:      thing.Price,
		Status:     ColumnQuoteStatusToStorage(thing.Status),
	}

}

// 数据库模型转换为领域模型
func ConvertColumnQuoteDOs(things []*model.ColumnQuote) []*db_model.ColumnQuoteDO {
	if things == nil {
		return nil
	}
	var readerDOs []*db_model.ColumnQuoteDO
	for _, item := range things {
		readerDOs = append(readerDOs, ConvertColumnQuoteDO(item))
	}
	return readerDOs
}
