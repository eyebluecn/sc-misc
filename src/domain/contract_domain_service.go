package domain

import (
	"context"
	"fmt"
	"github.com/eyebluecn/sc-misc/src/common/enums"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/infra/mq"
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
	"github.com/eyebluecn/sc-misc/src/util"
)

type ContractDomainService struct{}

func NewContractDomainService() *ContractDomainService {
	return &ContractDomainService{}
}

// 新增用户
func (receiver ContractDomainService) Create(ctx context.Context, column *model.Column, author *model.Author) (*model.Contract, error) {

	//参数校验
	err := receiver.createParamCheck(ctx, column, author)
	if err != nil {
		return nil, err
	}

	contract := &model.Contract{
		Name:       fmt.Sprintf("《%v》合同", column.Name),
		Content:    "这里是合同的正文内容",
		ColumnID:   column.ID,
		AuthorID:   author.ID,
		Status:     model.ContractStatusOk,
		Percentage: 0.4,
		PaymentDay: "LAST_DAY_OF_MONTH",
	}

	contractRepo := repo.NewContractRepo()
	contract, err = contractRepo.Insert(ctx, contract)
	if err != nil {
		return nil, err
	}

	//发出领域事件
	_ = mq.NewProducer().Publish(ctx, mq.MqTopicContract, "contract", "CONTRACT_CREATE", util.ToJSON(column))

	return contract, nil
}

func (receiver ContractDomainService) createParamCheck(ctx context.Context, column *model.Column, author *model.Author) error {

	//参数校验。
	if column == nil {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "column is nil")
	}
	if author == nil {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "author is nil")
	}
	if column.Name == "" {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "用户名不能为空")
	}
	if column.AuthorID != author.ID {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "作者信息不一致")
	}

	return nil

}
