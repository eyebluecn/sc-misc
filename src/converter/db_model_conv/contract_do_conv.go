package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
)

// 枚举转为存储的整数
func ContractStatusToStorage(status model.ContractStatus) int32 {
	return int32(status)
}

// 数据库模型转换为领域模型
func ConvertContractDO(thing *model.Contract) *db_model.ContractDO {
	if thing == nil {
		return nil
	}

	return &db_model.ContractDO{
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
func ConvertContractDOs(things []*model.Contract) []*db_model.ContractDO {
	if things == nil {
		return nil
	}
	var results []*db_model.ContractDO
	for _, item := range things {
		results = append(results, ConvertContractDO(item))
	}
	return results
}
