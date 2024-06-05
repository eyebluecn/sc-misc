package enums

type SubscriptionStatus int32

const (
	SubscriptionStatusCreated  SubscriptionStatus = 0 //已创建
	SubscriptionStatusOk       SubscriptionStatus = 1 //已生效
	SubscriptionStatusDisabled SubscriptionStatus = 2 //已失效
)
