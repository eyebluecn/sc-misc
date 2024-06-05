package do2po

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 枚举转为存储的整数
func ConvertColumnStatus(status enums.ColumnStatus) int32 {
	return int32(status)
}

// 数据库模型转换为领域模型
func ConvertColumnPO(thing *do.ColumnDO) *po.ColumnPO {
	if thing == nil {
		return nil
	}

	return &po.ColumnPO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Name:       thing.Name,
		AuthorID:   thing.AuthorID,
		Status:     ConvertColumnStatus(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertColumnPOs(things []*do.ColumnDO) []*po.ColumnPO {
	if things == nil {
		return nil
	}
	var results []*po.ColumnPO
	for _, item := range things {
		results = append(results, ConvertColumnPO(item))
	}
	return results
}
