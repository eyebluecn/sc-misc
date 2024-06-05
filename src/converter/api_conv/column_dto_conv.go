package api_conv

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
)

// 转为枚举
func ConvertColumnStatus(status enums.ColumnStatus) sc_misc_api.ColumnStatus {
	return sc_misc_api.ColumnStatus(status)
}

// 领域模型转为传输模型
func ConvertColumnDTO(thing *do.Column) *sc_misc_api.ColumnDTO {
	if thing == nil {
		return nil
	}

	return &sc_misc_api.ColumnDTO{
		Id:         thing.ID,
		CreateTime: util.Timestamp(thing.CreateTime),
		UpdateTime: util.Timestamp(thing.UpdateTime),
		Name:       thing.Name,
		AuthorId:   thing.AuthorID,
		Status:     ConvertColumnStatus(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertColumnDTOs(things []*do.Column) []*sc_misc_api.ColumnDTO {
	if things == nil {
		return nil
	}
	var readerDTOs []*sc_misc_api.ColumnDTO
	for _, item := range things {
		readerDTOs = append(readerDTOs, ConvertColumnDTO(item))
	}
	return readerDTOs
}
