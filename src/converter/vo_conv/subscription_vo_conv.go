package vo_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/vo_model"
	"github.com/eyebluecn/sc-misc/src/util"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// 转为枚举
func ConvertSubscriptionStatus(status sc_subscription_api.SubscriptionStatus) vo_model.SubscriptionStatus {
	return vo_model.SubscriptionStatus(status)
}

// 转为VO
func ConvertSubscriptionVO(thing *sc_subscription_api.SubscriptionDTO) *vo_model.SubscriptionVO {
	if thing == nil {
		return nil
	}
	return &vo_model.SubscriptionVO{
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
func ConvertSubscriptions(things []*sc_subscription_api.SubscriptionDTO) []*vo_model.SubscriptionVO {
	if things == nil {
		return nil
	}
	var subscriptionVOS []*vo_model.SubscriptionVO
	for _, item := range things {
		subscriptionVOS = append(subscriptionVOS, ConvertSubscriptionVO(item))
	}
	return subscriptionVOS
}
