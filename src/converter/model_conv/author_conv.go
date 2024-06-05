package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 数据库模型转换为领域模型
func ConvertAuthor(thing *po.AuthorPO) *do.Author {
	if thing == nil {
		return nil
	}

	return &do.Author{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Username:   thing.Username,
		Password:   thing.Password,
		Realname:   thing.Realname,
	}
}

// 数据库模型转换为领域模型
func ConvertAuthors(things []*po.AuthorPO) []*do.Author {
	if things == nil {
		return nil
	}
	var readers []*do.Author
	for _, item := range things {
		readers = append(readers, ConvertAuthor(item))
	}
	return readers
}
