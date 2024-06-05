package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 数据库模型转换为领域模型
func ConvertReaderDO(thing *do.Reader) *po.ReaderPO {
	if thing == nil {
		return nil
	}

	return &po.ReaderPO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Username:   thing.Username,
		Password:   thing.Password,
	}
}

// 数据库模型转换为领域模型
func ConvertReaderDOs(things []*do.Reader) []*po.ReaderPO {
	if things == nil {
		return nil
	}
	var readerDOs []*po.ReaderPO
	for _, item := range things {
		readerDOs = append(readerDOs, ConvertReaderDO(item))
	}
	return readerDOs
}
