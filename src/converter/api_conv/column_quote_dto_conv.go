package api_conv

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
)

// 转为枚举
func ConvertColumnQuoteStatus(status enums.ColumnQuoteStatus) sc_misc_api.ColumnQuoteStatus {
	return sc_misc_api.ColumnQuoteStatus(status)
}

// 领域模型转为传输模型
func ConvertColumnQuoteDTO(thing *do.ColumnQuoteDO) *sc_misc_api.ColumnQuoteDTO {
	if thing == nil {
		return nil
	}

	return &sc_misc_api.ColumnQuoteDTO{
		Id:         thing.ID,
		CreateTime: util.Timestamp(thing.CreateTime),
		UpdateTime: util.Timestamp(thing.UpdateTime),
		ColumnId:   thing.ColumnID,
		EditorId:   thing.EditorID,
		Price:      thing.Price,
		Status:     ConvertColumnQuoteStatus(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertColumnQuoteDTOs(things []*do.ColumnQuoteDO) []*sc_misc_api.ColumnQuoteDTO {
	if things == nil {
		return nil
	}
	var readerDTOs []*sc_misc_api.ColumnQuoteDTO
	for _, item := range things {
		readerDTOs = append(readerDTOs, ConvertColumnQuoteDTO(item))
	}
	return readerDTOs
}
