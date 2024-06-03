package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
)

// 数据库模型转换为领域模型
func ConvertReaderDO(thing *model.Reader) *db_model.ReaderDO {
	if thing == nil {
		return nil
	}

	return &db_model.ReaderDO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Username:   thing.Username,
		Password:   thing.Password,
	}
}

// 数据库模型转换为领域模型
func ConvertReaderDOs(things []*model.Reader) []*db_model.ReaderDO {
	if things == nil {
		return nil
	}
	var readerDOs []*db_model.ReaderDO
	for _, item := range things {
		readerDOs = append(readerDOs, ConvertReaderDO(item))
	}
	return readerDOs
}
