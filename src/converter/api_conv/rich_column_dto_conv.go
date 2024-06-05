package api_conv

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/model/info"
)

// 领域模型转为传输模型
func ConvertRichColumnDTO(thing *info.RichColumn) *sc_misc_api.RichColumnDTO {
	if thing == nil {
		return nil
	}

	return &sc_misc_api.RichColumnDTO{
		Column:       ConvertColumnDTO(thing.Column),
		Author:       ConvertAuthorDTO(thing.Author),
		ColumnQuote:  ConvertColumnQuoteDTO(thing.ColumnQuote),
		Subscription: ConvertSubscriptionDTO(thing.Subscription),
	}
}

// 数据库模型转换为领域模型
func ConvertRichColumnDTOs(things []*info.RichColumn) []*sc_misc_api.RichColumnDTO {
	if things == nil {
		return nil
	}
	var readerDTOs []*sc_misc_api.RichColumnDTO
	for _, item := range things {
		readerDTOs = append(readerDTOs, ConvertRichColumnDTO(item))
	}
	return readerDTOs
}
