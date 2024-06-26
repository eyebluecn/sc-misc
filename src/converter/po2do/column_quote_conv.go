package po2do

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
func ConvertColumnQuoteDO(thing *po.ColumnQuotePO) *do.ColumnQuoteDO {
	if thing == nil {
		return nil
	}

	return &do.ColumnQuoteDO{
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
func ConvertColumnQuoteDOs(things []*po.ColumnQuotePO) []*do.ColumnQuoteDO {
	if things == nil {
		return nil
	}
	var readers []*do.ColumnQuoteDO
	for _, item := range things {
		readers = append(readers, ConvertColumnQuoteDO(item))
	}
	return readers
}
