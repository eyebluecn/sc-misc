package do2dto

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/model/do"
)

// 领域模型转为传输模型
func ConvertReaderDTO(thing *do.ReaderDO) *sc_misc_api.ReaderDTO {
	if thing == nil {
		return nil
	}

	return &sc_misc_api.ReaderDTO{
		Id:         thing.ID,
		CreateTime: util.Timestamp(thing.CreateTime),
		UpdateTime: util.Timestamp(thing.UpdateTime),
		Username:   thing.Username,
	}
}

// 数据库模型转换为领域模型
func ConvertReaderDTOs(things []*do.ReaderDO) []*sc_misc_api.ReaderDTO {
	if things == nil {
		return nil
	}
	var readerDTOs []*sc_misc_api.ReaderDTO
	for _, item := range things {
		readerDTOs = append(readerDTOs, ConvertReaderDTO(item))
	}
	return readerDTOs
}
