package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
)

// 枚举转为存储的整数
func PaymentStatusToStorage(status model.PaymentStatus) int32 {
	return int32(status)
}

// 数据库模型转换为领域模型
func ConvertPaymentDO(thing *model.Payment) *db_model.PaymentDO {
	if thing == nil {
		return nil
	}

	return &db_model.PaymentDO{
		ID:                 thing.ID,
		CreateTime:         thing.CreateTime,
		UpdateTime:         thing.UpdateTime,
		OrderNo:            thing.OrderNo,
		Method:             thing.Method,
		ThirdTransactionNo: thing.ThirdTransactionNo,
		Amount:             thing.Amount,
		Status:             PaymentStatusToStorage(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertPaymentDOs(things []*model.Payment) []*db_model.PaymentDO {
	if things == nil {
		return nil
	}
	var results []*db_model.PaymentDO
	for _, item := range things {
		results = append(results, ConvertPaymentDO(item))
	}
	return results
}
