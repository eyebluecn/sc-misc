package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/universal"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_base"
)

// 转为分页
func ConvertSubscriptionPagination(pagination *sc_subscription_base.Pagination) *universal.Pagination {
	if pagination == nil {
		return nil
	}
	return &universal.Pagination{
		PageNum:    pagination.PageNum,
		PageSize:   pagination.PageSize,
		TotalItems: pagination.TotalItems,
	}
}
