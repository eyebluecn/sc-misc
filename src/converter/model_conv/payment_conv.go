package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
)

// 转为枚举
func ConvertPaymentStatus(status int32) model.PaymentStatus {
	return model.PaymentStatus(status)
}

// 数据库模型转换为领域模型
func ConvertPayment(thing *db_model.PaymentDO) *model.Payment {
	if thing == nil {
		return nil
	}

	return &model.Payment{
		ID:                 thing.ID,
		CreateTime:         thing.CreateTime,
		UpdateTime:         thing.UpdateTime,
		OrderNo:            thing.OrderNo,
		Method:             thing.Method,
		ThirdTransactionNo: thing.ThirdTransactionNo,
		Amount:             thing.Amount,
		Status:             ConvertPaymentStatus(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertPayments(things []*db_model.PaymentDO) []*model.Payment {
	if things == nil {
		return nil
	}
	var readers []*model.Payment
	for _, item := range things {
		readers = append(readers, ConvertPayment(item))
	}
	return readers
}
