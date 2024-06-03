package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
)

// 转为枚举
func ConvertColumnStatus(status int32) model.ColumnStatus {
	return model.ColumnStatus(status)
}

// 数据库模型转换为领域模型
func ConvertColumn(thing *db_model.ColumnDO) *model.Column {
	if thing == nil {
		return nil
	}

	return &model.Column{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Name:       thing.Name,
		AuthorID:   thing.AuthorID,
		Status:     ConvertColumnStatus(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertColumns(things []*db_model.ColumnDO) []*model.Column {
	if things == nil {
		return nil
	}
	var readers []*model.Column
	for _, item := range things {
		readers = append(readers, ConvertColumn(item))
	}
	return readers
}
