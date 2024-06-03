package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
)

// 数据库模型转换为领域模型
func ConvertAuthorDO(thing *model.Author) *db_model.AuthorDO {
	if thing == nil {
		return nil
	}

	return &db_model.AuthorDO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Username:   thing.Username,
		Password:   thing.Password,
		Realname:   thing.Realname,
	}
}

// 数据库模型转换为领域模型
func ConvertAuthorDOs(things []*model.Author) []*db_model.AuthorDO {
	if things == nil {
		return nil
	}
	var readerDOs []*db_model.AuthorDO
	for _, item := range things {
		readerDOs = append(readerDOs, ConvertAuthorDO(item))
	}
	return readerDOs
}
