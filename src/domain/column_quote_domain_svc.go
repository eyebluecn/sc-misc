package domain

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/common/util"
	mq2 "github.com/eyebluecn/sc-misc/src/infrastructure/middleware/mq"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type ColumnQuoteDomainSvc struct{}

func NewColumnQuoteDomainSvc() *ColumnQuoteDomainSvc {
	return &ColumnQuoteDomainSvc{}
}

// 新增用户
func (receiver ColumnQuoteDomainSvc) Create(ctx context.Context, column *do.ColumnDO, price int64, editor *do.EditorDO) (*do.ColumnQuoteDO, error) {

	//参数校验
	err := receiver.createParamCheck(ctx, column, editor)
	if err != nil {
		return nil, err
	}

	contract := &do.ColumnQuoteDO{
		ColumnID: column.ID,
		EditorID: editor.ID,
		Price:    price,
		Status:   enums.ColumnQuoteStatusOk,
	}

	contractRepo := repo.NewColumnQuoteRepo()
	contract, err = contractRepo.Insert(ctx, contract)
	if err != nil {
		return nil, err
	}

	//发出领域事件
	_ = mq2.DefaultProducer().Publish(ctx, mq2.MqTopicContract, "contract", "CONTRACT_CREATE", util.ToJSON(column))

	return contract, nil
}

func (receiver ColumnQuoteDomainSvc) createParamCheck(ctx context.Context, column *do.ColumnDO, editor *do.EditorDO) error {

	//参数校验。
	if column == nil {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "column is nil")
	}
	if editor == nil {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "author is nil")
	}
	if column.Name == "" {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "用户名不能为空")
	}

	return nil

}
