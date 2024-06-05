package dto2result

import (
	"github.com/eyebluecn/sc-misc/src/model/result"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_base"
)

// 转为分页
func ConvertSubscriptionPagination(pagination *sc_subscription_base.Pagination) *result.Pagination {
	if pagination == nil {
		return nil
	}
	return &result.Pagination{
		PageNum:    pagination.PageNum,
		PageSize:   pagination.PageSize,
		TotalItems: pagination.TotalItems,
	}
}
