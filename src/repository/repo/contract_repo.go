package repo

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/common/config"
	"github.com/eyebluecn/sc-misc/src/converter/db_model_conv"
	"github.com/eyebluecn/sc-misc/src/converter/model_conv"
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/query"
	"time"
)

type ContractRepo struct {
}

func NewContractRepo() ContractRepo {
	return ContractRepo{}
}

// 新建一个Reader
func (receiver ContractRepo) Insert(
	ctx context.Context,
	contract *model.Contract,
) (*model.Contract, error) {
	table := query.Use(config.DB).ContractDO

	//时间置为当前
	contract.CreateTime = time.Now()
	contract.UpdateTime = time.Now()

	contractDO := db_model_conv.ConvertContractDO(contract)

	err := table.WithContext(ctx).Debug().Create(contractDO)
	if err != nil {
		return nil, err
	}

	return model_conv.ConvertContract(contractDO), nil
}
