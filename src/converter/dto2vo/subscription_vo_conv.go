package dto2vo

import (
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/model/vo"
	"github.com/eyebluecn/sc-misc/src/model/vo/enums"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// 转为枚举
func ConvertSubscriptionStatus(status sc_subscription_api.SubscriptionStatus) enums.SubscriptionStatus {
	return enums.SubscriptionStatus(status)
}

// 转为VO
func ConvertSubscriptionVO(thing *sc_subscription_api.SubscriptionDTO) *vo.SubscriptionVO {
	if thing == nil {
		return nil
	}
	return &vo.SubscriptionVO{
		ID:         thing.Id,
		CreateTime: util.ParseTimestamp(thing.CreateTime),
		UpdateTime: util.ParseTimestamp(thing.UpdateTime),
		ReaderID:   thing.ReaderId,
		ColumnID:   thing.ColumnId,
		OrderID:    thing.OrderId,
		Status:     ConvertSubscriptionStatus(thing.Status),
	}
}

// 转为VO数组
func ConvertSubscriptionVOs(things []*sc_subscription_api.SubscriptionDTO) []*vo.SubscriptionVO {
	if things == nil {
		return nil
	}
	var subscriptionVOS []*vo.SubscriptionVO
	for _, item := range things {
		subscriptionVOS = append(subscriptionVOS, ConvertSubscriptionVO(item))
	}
	return subscriptionVOS
}
