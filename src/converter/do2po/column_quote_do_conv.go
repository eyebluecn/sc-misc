package do2po

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 枚举转为存储的整数
func ConvertColumnQuoteStatus(status enums.ColumnQuoteStatus) int32 {
	return int32(status)
}

// 数据库模型转换为领域模型
func ConvertColumnQuotePO(thing *do.ColumnQuoteDO) *po.ColumnQuotePO {
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
		Status:     ConvertColumnQuoteStatus(thing.Status),
	}

}

// 数据库模型转换为领域模型
func ConvertColumnQuotePOs(things []*do.ColumnQuoteDO) []*po.ColumnQuotePO {
	if things == nil {
		return nil
	}
	var readerDOs []*po.ColumnQuotePO
	for _, item := range things {
		readerDOs = append(readerDOs, ConvertColumnQuotePO(item))
	}
	return readerDOs
}
