package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
)

// 数据库模型转换为领域模型
func ConvertReader(thing *db_model.ReaderDO) *model.Reader {
	if thing == nil {
		return nil
	}

	return &model.Reader{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Username:   thing.Username,
		Password:   thing.Password,
	}
}

// 数据库模型转换为领域模型
func ConvertReaders(things []*db_model.ReaderDO) []*model.Reader {
	if things == nil {
		return nil
	}
	var readers []*model.Reader
	for _, item := range things {
		readers = append(readers, ConvertReader(item))
	}
	return readers
}
