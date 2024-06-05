package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 数据库模型转换为领域模型
func ConvertReader(thing *po.ReaderPO) *do.ReaderDO {
	if thing == nil {
		return nil
	}

	return &do.ReaderDO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Username:   thing.Username,
		Password:   thing.Password,
	}
}

// 数据库模型转换为领域模型
func ConvertReaders(things []*po.ReaderPO) []*do.ReaderDO {
	if things == nil {
		return nil
	}
	var readers []*do.ReaderDO
	for _, item := range things {
		readers = append(readers, ConvertReader(item))
	}
	return readers
}
