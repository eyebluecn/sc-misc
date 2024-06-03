package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
)

// 转为枚举
func ConvertContractStatus(status int32) model.ContractStatus {
	return model.ContractStatus(status)
}

// 数据库模型转换为领域模型
func ConvertContract(thing *db_model.ContractDO) *model.Contract {
	if thing == nil {
		return nil
	}

	return &model.Contract{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Name:       thing.Name,
		Content:    thing.Content,
		ColumnID:   thing.ColumnID,
		AuthorID:   thing.AuthorID,
		Status:     ConvertContractStatus(thing.Status),
		Percentage: thing.Percentage,
		PaymentDay: thing.PaymentDay,
	}
}

// 数据库模型转换为领域模型
func ConvertContracts(things []*db_model.ContractDO) []*model.Contract {
	if things == nil {
		return nil
	}
	var results []*model.Contract
	for _, item := range things {
		results = append(results, ConvertContract(item))
	}
	return results
}
