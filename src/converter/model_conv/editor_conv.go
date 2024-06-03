package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
)

// 数据库模型转换为领域模型
func ConvertEditor(thing *db_model.EditorDO) *model.Editor {
	if thing == nil {
		return nil
	}

	return &model.Editor{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Username:   thing.Username,
		Password:   thing.Password,
		WorkNo:     thing.WorkNo,
	}
}

// 数据库模型转换为领域模型
func ConvertEditors(things []*db_model.EditorDO) []*model.Editor {
	if things == nil {
		return nil
	}
	var readers []*model.Editor
	for _, item := range things {
		readers = append(readers, ConvertEditor(item))
	}
	return readers
}
