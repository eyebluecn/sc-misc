package model

import "github.com/eyebluecn/sc-misc/src/model/vo_model"

// 专栏领域模型  这个只是视图模型。
type RichColumn struct {
	Column       *Column
	Author       *Author
	ColumnQuote  *ColumnQuote
	Subscription *vo_model.SubscriptionVO
}
