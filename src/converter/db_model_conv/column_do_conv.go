package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 枚举转为存储的整数
func ColumnStatusToStorage(status enums.ColumnStatus) int32 {
	return int32(status)
}

// 数据库模型转换为领域模型
func ConvertColumnDO(thing *do.Column) *po.ColumnPO {
	if thing == nil {
		return nil
	}

	return &po.ColumnPO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Name:       thing.Name,
		AuthorID:   thing.AuthorID,
		Status:     ColumnStatusToStorage(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertColumnDOs(things []*do.Column) []*po.ColumnPO {
	if things == nil {
		return nil
	}
	var results []*po.ColumnPO
	for _, item := range things {
		results = append(results, ConvertColumnDO(item))
	}
	return results
}
