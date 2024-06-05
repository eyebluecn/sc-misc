package info

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/vo"
)

// 专栏领域模型  这个只是视图模型。
type RichColumn struct {
	Column       *do.ColumnDO
	Author       *do.AuthorDO
	ColumnQuote  *do.ColumnQuoteDO
	Subscription *vo.SubscriptionVO
}
