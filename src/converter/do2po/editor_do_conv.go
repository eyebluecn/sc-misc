package do2po

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

func ConvertEditorPO(thing *do.EditorDO) *po.EditorPO {
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

func ConvertEditorPOs(things []*do.EditorDO) []*po.EditorPO {
	if things == nil {
		return nil
	}
	var readerDOs []*po.EditorPO
	for _, item := range things {
		readerDOs = append(readerDOs, ConvertEditorPO(item))
	}
	return readerDOs
}
