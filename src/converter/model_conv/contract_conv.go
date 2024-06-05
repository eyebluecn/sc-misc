package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 转为枚举
func ConvertContractStatus(status int32) enums.ContractStatus {
	return enums.ContractStatus(status)
}

// 数据库模型转换为领域模型
func ConvertContract(thing *po.ContractPO) *do.Contract {
	if thing == nil {
		return nil
	}

	return &do.Contract{
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
func ConvertContracts(things []*po.ContractPO) []*do.Contract {
	if things == nil {
		return nil
	}
	var results []*do.Contract
	for _, item := range things {
		results = append(results, ConvertContract(item))
	}
	return results
}
