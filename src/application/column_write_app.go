package application

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/domain"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/info"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type ColumnWriteApp struct{}

func NewColumnWriteApp() *ColumnWriteApp {
	return &ColumnWriteApp{}
}

/**
 * ！！！此接口仅作为演示用途！！！
 * 一口气创建 作者，作者合同，专栏，课程文章，专栏报价，编辑
 * 为了保证数据库不出现脏数据
 */
func (receiver ColumnWriteApp) ColumnOmnibus(ctx context.Context, request sc_misc_api.ColumnOmnibusRequest) (*info.RichColumn, error) {

	//手动开启事务
	//根据作者名，寻找作者。
	author, err := repo.NewAuthorRepo().FindByUsername(ctx, request.AuthorName)
	if err != nil {
		return nil, err
	}
	if author == nil {
		//作者注册，简单用密码123456，真名和用户名一样。
		author = &do.AuthorDO{
			Username: request.AuthorName,
			Password: "123456",
			Realname: request.AuthorName,
		}
		author, err = domain.NewAuthorDomainService().Create(ctx, author)
		if err != nil {
			return nil, err
		}
	}

	//查找小编
	editor, err := repo.NewEditorRepo().CheckById(ctx, request.Operator.OperatorId)
	if err != nil {
		return nil, err
	}

	//创建专栏
	column := &do.ColumnDO{
		Name:     request.ColumnName,
		AuthorID: author.ID,
		Status:   enums.ColumnStatusOk,
	}
	column, err = domain.NewColumnDomainService().Create(ctx, column, author)
	if err != nil {
		return nil, err
	}

	//创建作者合同
	contract, err := domain.NewContractDomainService().Create(ctx, column, author)
	if err != nil {
		return nil, err
	}

	//创建专栏保价
	columnQuote, err := domain.NewColumnQuoteDomainService().Create(ctx, column, request.ColumnPrice, editor)
	if err != nil {
		return nil, err
	}

	klog.CtxInfof(ctx, "完成了混合创建：authorId=%v editorId=%v columnId=%v contractId=%v columnQuoteId=%v",
		author.ID, editor.ID, column.ID, contract.ID, columnQuote.ID)

	richColumn := &info.RichColumn{
		Column:       column,
		Author:       author,
		ColumnQuote:  columnQuote,
		Subscription: nil,
	}

	return richColumn, nil
}
