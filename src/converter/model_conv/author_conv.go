package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
)

// 数据库模型转换为领域模型
func ConvertAuthor(thing *db_model.AuthorDO) *model.Author {
	if thing == nil {
		return nil
	}

	return &model.Author{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Username:   thing.Username,
		Password:   thing.Password,
		Realname:   thing.Realname,
	}
}

// 数据库模型转换为领域模型
func ConvertAuthors(things []*db_model.AuthorDO) []*model.Author {
	if things == nil {
		return nil
	}
	var readers []*model.Author
	for _, item := range things {
		readers = append(readers, ConvertAuthor(item))
	}
	return readers
}
