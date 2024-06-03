package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
)

// 数据库模型转换为领域模型
func ConvertEditorDO(thing *model.Editor) *db_model.EditorDO {
	if thing == nil {
		return nil
	}

	return &db_model.EditorDO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Username:   thing.Username,
		Password:   thing.Password,
		WorkNo:     thing.WorkNo,
	}
}

// 数据库模型转换为领域模型
func ConvertEditorDOs(things []*model.Editor) []*db_model.EditorDO {
	if things == nil {
		return nil
	}
	var readerDOs []*db_model.EditorDO
	for _, item := range things {
		readerDOs = append(readerDOs, ConvertEditorDO(item))
	}
	return readerDOs
}
