package po2do

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 数据库模型转换为领域模型
func ConvertAuthorDO(thing *po.AuthorPO) *do.AuthorDO {
	if thing == nil {
		return nil
	}

	return &do.AuthorDO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Username:   thing.Username,
		Password:   thing.Password,
		Realname:   thing.Realname,
	}
}

// 数据库模型转换为领域模型
func ConvertAuthorDOs(things []*po.AuthorPO) []*do.AuthorDO {
	if things == nil {
		return nil
	}
	var readers []*do.AuthorDO
	for _, item := range things {
		readers = append(readers, ConvertAuthorDO(item))
	}
	return readers
}
