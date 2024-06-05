package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 枚举转为存储的整数
func ColumnQuoteStatusToStorage(status enums.ColumnQuoteStatus) int32 {
	return int32(status)
}

// 数据库模型转换为领域模型
func ConvertColumnQuoteDO(thing *do.ColumnQuoteDO) *po.ColumnQuotePO {
	if thing == nil {
		return nil
	}

	return &po.ColumnQuotePO{
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
func ConvertColumnQuoteDOs(things []*do.ColumnQuoteDO) []*po.ColumnQuotePO {
	if things == nil {
		return nil
	}
	var readerDOs []*po.ColumnQuotePO
	for _, item := range things {
		readerDOs = append(readerDOs, ConvertColumnQuoteDO(item))
	}
	return readerDOs
}
