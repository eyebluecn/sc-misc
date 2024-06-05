package model

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/vo"
)

// 专栏领域模型  这个只是视图模型。
type RichColumn struct {
	Column       *do.Column
	Author       *do.Author
	ColumnQuote  *do.ColumnQuote
	Subscription *vo.SubscriptionVO
}
