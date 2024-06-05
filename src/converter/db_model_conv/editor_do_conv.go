package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 数据库模型转换为领域模型
func ConvertEditorDO(thing *do.EditorDO) *po.EditorPO {
	if thing == nil {
		return nil
	}

	return &po.EditorPO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Username:   thing.Username,
		Password:   thing.Password,
		WorkNo:     thing.WorkNo,
	}
}

// 数据库模型转换为领域模型
func ConvertEditorDOs(things []*do.EditorDO) []*po.EditorPO {
	if things == nil {
		return nil
	}
	var readerDOs []*po.EditorPO
	for _, item := range things {
		readerDOs = append(readerDOs, ConvertEditorDO(item))
	}
	return readerDOs
}
