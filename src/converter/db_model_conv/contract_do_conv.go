package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 枚举转为存储的整数
func ContractStatusToStorage(status enums.ContractStatus) int32 {
	return int32(status)
}

// 数据库模型转换为领域模型
func ConvertContractDO(thing *do.Contract) *po.ContractPO {
	if thing == nil {
		return nil
	}

	return &po.ContractPO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		Name:       thing.Name,
		Content:    thing.Content,
		ColumnID:   thing.ColumnID,
		AuthorID:   thing.AuthorID,
		Status:     ContractStatusToStorage(thing.Status),
		Percentage: thing.Percentage,
		PaymentDay: thing.PaymentDay,
	}
}

// 数据库模型转换为领域模型
func ConvertContractDOs(things []*do.Contract) []*po.ContractPO {
	if things == nil {
		return nil
	}
	var results []*po.ContractPO
	for _, item := range things {
		results = append(results, ConvertContractDO(item))
	}
	return results
}
