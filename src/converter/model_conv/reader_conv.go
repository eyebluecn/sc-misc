package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 数据库模型转换为领域模型
func ConvertReader(thing *po.ReaderPO) *do.Reader {
	if thing == nil {
		return nil
	}

	return &do.Reader{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Username:   thing.Username,
		Password:   thing.Password,
	}
}

// 数据库模型转换为领域模型
func ConvertReaders(things []*po.ReaderPO) []*do.Reader {
	if things == nil {
		return nil
	}
	var readers []*do.Reader
	for _, item := range things {
		readers = append(readers, ConvertReader(item))
	}
	return readers
}
