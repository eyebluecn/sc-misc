package domain

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/common/util"
	mq2 "github.com/eyebluecn/sc-misc/src/infrastructure/middleware/mq"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type ColumnDomainService struct{}

func NewColumnDomainService() *ColumnDomainService {
	return &ColumnDomainService{}
}

// 新增用户
func (receiver ColumnDomainService) Create(ctx context.Context, column *do.ColumnDO, author *do.AuthorDO) (*do.ColumnDO, error) {

	//参数校验
	err := receiver.createParamCheck(ctx, column, author)
	if err != nil {
		return nil, err
	}

	columnRepo := repo.NewColumnRepo()
	column, err = columnRepo.Insert(ctx, column)
	if err != nil {
		return nil, err
	}

	//发出领域事件
	_ = mq2.NewProducer().Publish(ctx, mq2.MqTopicColumn, "column", "COLUMN_CREATE", util.ToJSON(column))

	return column, nil
}

func (receiver ColumnDomainService) createParamCheck(ctx context.Context, column *do.ColumnDO, author *do.AuthorDO) error {

	//参数校验。
	if column == nil {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "column is nil")
	}
	if author == nil {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "author is nil")
	}
	if column.Name == "" {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "用户名不能为空")
	}
	if column.AuthorID != author.ID {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "作者信息不一致")
	}

	return nil

}
