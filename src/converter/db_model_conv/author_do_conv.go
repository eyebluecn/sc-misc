package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 数据库模型转换为领域模型
func ConvertAuthorDO(thing *do.Author) *po.AuthorPO {
	if thing == nil {
		return nil
	}

	return &po.AuthorPO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Username:   thing.Username,
		Password:   thing.Password,
		Realname:   thing.Realname,
	}
}

// 数据库模型转换为领域模型
func ConvertAuthorDOs(things []*do.Author) []*po.AuthorPO {
	if things == nil {
		return nil
	}
	var readerDOs []*po.AuthorPO
	for _, item := range things {
		readerDOs = append(readerDOs, ConvertAuthorDO(item))
	}
	return readerDOs
}
