package po2do

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 数据库模型转换为领域模型
func ConvertEditorDO(thing *po.EditorPO) *do.EditorDO {
	if thing == nil {
		return nil
	}

	return &do.EditorDO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Username:   thing.Username,
		Password:   thing.Password,
		WorkNo:     thing.WorkNo,
	}
}

// 数据库模型转换为领域模型
func ConvertEditorDOs(things []*po.EditorPO) []*do.EditorDO {
	if things == nil {
		return nil
	}
	var readers []*do.EditorDO
	for _, item := range things {
		readers = append(readers, ConvertEditorDO(item))
	}
	return readers
}
