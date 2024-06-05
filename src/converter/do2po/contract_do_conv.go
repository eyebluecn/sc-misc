package do2po

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 枚举转为存储的整数
func ConvertContractStatus(status enums.ContractStatus) int32 {
	return int32(status)
}

func ConvertContractPO(thing *do.ContractDO) *po.ContractPO {
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
		Status:     ConvertContractStatus(thing.Status),
		Percentage: thing.Percentage,
		PaymentDay: thing.PaymentDay,
	}
}

func ConvertContractPOs(things []*do.ContractDO) []*po.ContractPO {
	if things == nil {
		return nil
	}
	var results []*po.ContractPO
	for _, item := range things {
		results = append(results, ConvertContractPO(item))
	}
	return results
}
