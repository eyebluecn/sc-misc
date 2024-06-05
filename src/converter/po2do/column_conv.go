package po2do

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 转为枚举
func ConvertColumnStatus(status int32) enums.ColumnStatus {
	return enums.ColumnStatus(status)
}

// 数据库模型转换为领域模型
func ConvertColumnDO(thing *po.ColumnPO) *do.ColumnDO {
	if thing == nil {
		return nil
	}

	return &do.ColumnDO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Name:       thing.Name,
		AuthorID:   thing.AuthorID,
		Status:     ConvertColumnStatus(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertColumnDOs(things []*po.ColumnPO) []*do.ColumnDO {
	if things == nil {
		return nil
	}
	var readers []*do.ColumnDO
	for _, item := range things {
		readers = append(readers, ConvertColumnDO(item))
	}
	return readers
}
