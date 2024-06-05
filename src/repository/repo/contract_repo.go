package repo

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/common/config"
	"github.com/eyebluecn/sc-misc/src/converter/db_model_conv"
	"github.com/eyebluecn/sc-misc/src/converter/model_conv"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/repository/dao"
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
	contract *do.ContractDO,
) (*do.ContractDO, error) {
	table := dao.Use(config.DB).ContractPO

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
