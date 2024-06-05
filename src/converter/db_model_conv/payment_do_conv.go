package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 枚举转为存储的整数
func PaymentStatusToStorage(status enums.PaymentStatus) int32 {
	return int32(status)
}

// 数据库模型转换为领域模型
func ConvertPaymentDO(thing *do.Payment) *po.PaymentPO {
	if thing == nil {
		return nil
	}

	return &po.PaymentPO{
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
func ConvertPaymentDOs(things []*do.Payment) []*po.PaymentPO {
	if things == nil {
		return nil
	}
	var results []*po.PaymentPO
	for _, item := range things {
		results = append(results, ConvertPaymentDO(item))
	}
	return results
}
