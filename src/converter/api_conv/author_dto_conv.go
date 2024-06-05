package api_conv

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/model/do"
)

// 领域模型转为传输模型
func ConvertAuthorDTO(thing *do.Author) *sc_misc_api.AuthorDTO {
	if thing == nil {
		return nil
	}

	return &sc_misc_api.AuthorDTO{
		Id:         thing.ID,
		CreateTime: util.Timestamp(thing.CreateTime),
		UpdateTime: util.Timestamp(thing.UpdateTime),
		Username:   thing.Username,
	}
}

// 数据库模型转换为领域模型
func ConvertAuthorDTOs(things []*do.Author) []*sc_misc_api.AuthorDTO {
	if things == nil {
		return nil
	}
	var readerDTOs []*sc_misc_api.AuthorDTO
	for _, item := range things {
		readerDTOs = append(readerDTOs, ConvertAuthorDTO(item))
	}
	return readerDTOs
}
