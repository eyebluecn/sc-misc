package application

import (
	"context"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_base"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/infrastructure/rpc"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/info"
	"github.com/eyebluecn/sc-misc/src/model/query"
	"github.com/eyebluecn/sc-misc/src/model/universal"
	"github.com/eyebluecn/sc-misc/src/model/vo"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type ColumnReadAppSvc struct{}

func NewColumnReadAppSvc() *ColumnReadAppSvc {
	return &ColumnReadAppSvc{}
}

// 获取某位读者查看到的专栏列表。
func (receiver ColumnReadAppSvc) ReaderColumnList(ctx context.Context, operator sc_misc_base.Operator, repoRequest query.ColumnPageQuery) ([]*info.RichColumn, *universal.Pagination, error) {

	columns, pagination, err := repo.NewColumnRepo().Page(ctx, repoRequest)
	if err != nil {
		return nil, nil, err
	}

	var richColumns []*info.RichColumn
	//装填专栏骨架
	for _, column := range columns {
		richColumns = append(richColumns, &info.RichColumn{
			Column: column,
		})
	}

	//依次装填. 这里可以做成并行填充。
	err = receiver.PopulateAuthor(ctx, richColumns)
	if err != nil {
		return nil, nil, err
	}

	err = receiver.PopulateColumnQuote(ctx, richColumns)
	if err != nil {
		return nil, nil, err
	}

	//如果是读者类型，那么需要填充订阅情况
	if operator.Type == sc_misc_base.OperatorType_READER {
		err = receiver.PopulateSubscription(ctx, richColumns, operator.OperatorId)
		if err != nil {
			return nil, nil, err
		}
	}

	return richColumns, pagination, nil
}

// 填充作者信息
func (receiver ColumnReadAppSvc) PopulateAuthor(ctx context.Context, richColumns []*info.RichColumn) error {

	var authorIds []int64
	for _, richColumn := range richColumns {
		authorIds = util.UniqueAddInt64(authorIds, richColumn.Column.AuthorID)
	}

	if len(authorIds) == 0 {
		return nil
	}

	list, err := repo.NewAuthorRepo().QueryByIds(ctx, authorIds)
	if err != nil {
		return err
	}

	//从list中找到合适的Author.
	for _, richColumn := range richColumns {
		author := receiver.findAuthor(ctx, list, richColumn.Column.AuthorID)
		if author != nil {
			richColumn.Author = author
		}
	}

	return nil

}

// 从列表中找到对应的作者
func (receiver ColumnReadAppSvc) findAuthor(ctx context.Context, authorList []*do.AuthorDO, authorId int64) *do.AuthorDO {
	for _, author := range authorList {
		if author.ID == authorId {
			return author
		}
	}
	return nil
}

// 填充价格信息
func (receiver ColumnReadAppSvc) PopulateColumnQuote(ctx context.Context, richColumns []*info.RichColumn) error {

	var columnIds []int64
	for _, richColumn := range richColumns {
		columnIds = util.UniqueAddInt64(columnIds, richColumn.Column.ID)
	}

	if len(columnIds) == 0 {
		return nil
	}

	list, err := repo.NewColumnQuoteRepo().FindOkByColumnIds(ctx, columnIds)
	if err != nil {
		return err
	}

	//从list中找到合适的ColumnQuote.
	for _, richColumn := range richColumns {
		author := receiver.findColumnQuote(ctx, list, richColumn.Column.ID)
		if author != nil {
			richColumn.ColumnQuote = author
		}
	}

	return nil

}

// 从列表中找到对应的作者
func (receiver ColumnReadAppSvc) findColumnQuote(ctx context.Context, columnQuotes []*do.ColumnQuoteDO, columnId int64) *do.ColumnQuoteDO {
	for _, columnQuote := range columnQuotes {
		if columnQuote.ID == columnId {
			return columnQuote
		}
	}
	return nil
}

// 填充订阅信息
func (receiver ColumnReadAppSvc) PopulateSubscription(ctx context.Context, richColumns []*info.RichColumn, readerId int64) error {

	var columnIds []int64
	for _, richColumn := range richColumns {
		columnIds = util.UniqueAddInt64(columnIds, richColumn.Column.ID)
	}

	if len(columnIds) == 0 {
		return nil
	}

	//查询读者对于columnIds的订阅情况。
	subscriptionList, err := rpc.NewSubscriptionCaller().SubscriptionList(ctx, readerId, columnIds)
	if err != nil {
		return err
	}

	//从list中找到合适的Subscription.
	for _, richColumn := range richColumns {
		author := receiver.findSubscription(ctx, subscriptionList, richColumn.Column.ID)
		if author != nil {
			richColumn.Subscription = author
		}
	}

	return nil

}

// 从列表中找到对应的作者
func (receiver ColumnReadAppSvc) findSubscription(ctx context.Context, subscriptionList []*vo.SubscriptionVO, columnId int64) *vo.SubscriptionVO {
	for _, subscription := range subscriptionList {
		if subscription.ColumnID == columnId {
			return subscription
		}
	}
	return nil
}

// 根据id来查询专栏信息
func (receiver ColumnReadAppSvc) QueryById(ctx context.Context, columnId int64) (*do.ColumnDO, error) {
	column, err := repo.NewColumnRepo().QueryById(ctx, columnId)
	if err != nil {
		return nil, err
	}
	return column, nil
}

// 根据id来查询专栏信息
func (receiver ColumnReadAppSvc) QueryByIds(ctx context.Context, columnIds []int64) ([]*do.ColumnDO, error) {
	columns, err := repo.NewColumnRepo().QueryByIds(ctx, columnIds)
	if err != nil {
		return nil, err
	}
	return columns, nil
}
