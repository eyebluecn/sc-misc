package do2dto

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_base"
	"github.com/eyebluecn/sc-misc/src/model/result"
)

// 领域模型转为传输模型
func ConvertPagination(thing *result.Pagination) *sc_misc_base.Pagination {
	if thing == nil {
		return nil
	}

	return &sc_misc_base.Pagination{
		PageNum:    thing.PageNum,
		PageSize:   thing.PageSize,
		TotalItems: thing.TotalItems,
	}
}
