package domain

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/infrastructure/middleware/mq"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type ColumnDomainSvc struct{}

func NewColumnDomainSvc() *ColumnDomainSvc {
	return &ColumnDomainSvc{}
}

// 新增用户
func (receiver ColumnDomainSvc) Create(ctx context.Context, column *do.ColumnDO, author *do.AuthorDO) (*do.ColumnDO, error) {

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
	_ = mq.DefaultProducer().Publish(ctx, mq.MqTopicColumn, "column", "COLUMN_CREATE", util.ToJSON(column))

	return column, nil
}

func (receiver ColumnDomainSvc) createParamCheck(ctx context.Context, column *do.ColumnDO, author *do.AuthorDO) error {

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
