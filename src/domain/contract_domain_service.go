package domain

import (
	"context"
	"fmt"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/common/util"
	mq2 "github.com/eyebluecn/sc-misc/src/infrastructure/middleware/mq"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type ContractDomainService struct{}

func NewContractDomainService() *ContractDomainService {
	return &ContractDomainService{}
}

// 新增用户
func (receiver ContractDomainService) Create(ctx context.Context, column *do.ColumnDO, author *do.AuthorDO) (*do.ContractDO, error) {

	//参数校验
	err := receiver.createParamCheck(ctx, column, author)
	if err != nil {
		return nil, err
	}

	contract := &do.ContractDO{
		Name:       fmt.Sprintf("《%v》合同", column.Name),
		Content:    "这里是合同的正文内容",
		ColumnID:   column.ID,
		AuthorID:   author.ID,
		Status:     enums.ContractStatusOk,
		Percentage: 0.4,
		PaymentDay: "LAST_DAY_OF_MONTH",
	}

	contractRepo := repo.NewContractRepo()
	contract, err = contractRepo.Insert(ctx, contract)
	if err != nil {
		return nil, err
	}

	//发出领域事件
	_ = mq2.NewProducer().Publish(ctx, mq2.MqTopicContract, "contract", "CONTRACT_CREATE", util.ToJSON(column))

	return contract, nil
}

func (receiver ContractDomainService) createParamCheck(ctx context.Context, column *do.ColumnDO, author *do.AuthorDO) error {

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
