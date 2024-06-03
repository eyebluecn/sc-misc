package domain

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/common/enums"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/infra/mq"
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
	"github.com/eyebluecn/sc-misc/src/util"
)

type ColumnQuoteDomainService struct{}

func NewColumnQuoteDomainService() *ColumnQuoteDomainService {
	return &ColumnQuoteDomainService{}
}

// 新增用户
func (receiver ColumnQuoteDomainService) Create(ctx context.Context, column *model.Column, price int64, editor *model.Editor) (*model.ColumnQuote, error) {

	//参数校验
	err := receiver.createParamCheck(ctx, column, editor)
	if err != nil {
		return nil, err
	}

	contract := &model.ColumnQuote{
		ColumnID: column.ID,
		EditorID: editor.ID,
		Price:    price,
		Status:   model.ColumnQuoteStatusOk,
	}

	contractRepo := repo.NewColumnQuoteRepo()
	contract, err = contractRepo.Insert(ctx, contract)
	if err != nil {
		return nil, err
	}

	//发出领域事件
	_ = mq.NewProducer().Publish(ctx, mq.MqTopicContract, "contract", "CONTRACT_CREATE", util.ToJSON(column))

	return contract, nil
}

func (receiver ColumnQuoteDomainService) createParamCheck(ctx context.Context, column *model.Column, editor *model.Editor) error {

	//参数校验。
	if column == nil {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "column is nil")
	}
	if editor == nil {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "author is nil")
	}
	if column.Name == "" {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "用户名不能为空")
	}

	return nil

}
