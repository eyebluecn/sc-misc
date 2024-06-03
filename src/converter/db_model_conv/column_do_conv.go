package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
)

// 枚举转为存储的整数
func ColumnStatusToStorage(status model.ColumnStatus) int32 {
	return int32(status)
}

// 数据库模型转换为领域模型
func ConvertColumnDO(thing *model.Column) *db_model.ColumnDO {
	if thing == nil {
		return nil
	}

	return &db_model.ColumnDO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Name:       thing.Name,
		AuthorID:   thing.AuthorID,
		Status:     ColumnStatusToStorage(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertColumnDOs(things []*model.Column) []*db_model.ColumnDO {
	if things == nil {
		return nil
	}
	var results []*db_model.ColumnDO
	for _, item := range things {
		results = append(results, ConvertColumnDO(item))
	}
	return results
}
