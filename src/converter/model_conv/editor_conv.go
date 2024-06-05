package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 数据库模型转换为领域模型
func ConvertEditor(thing *po.EditorPO) *do.Editor {
	if thing == nil {
		return nil
	}

	return &do.Editor{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Username:   thing.Username,
		Password:   thing.Password,
		WorkNo:     thing.WorkNo,
	}
}

// 数据库模型转换为领域模型
func ConvertEditors(things []*po.EditorPO) []*do.Editor {
	if things == nil {
		return nil
	}
	var readers []*do.Editor
	for _, item := range things {
		readers = append(readers, ConvertEditor(item))
	}
	return readers
}
