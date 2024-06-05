package po2do

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 转为枚举
func ConvertPaymentStatus(status int32) enums.PaymentStatus {
	return enums.PaymentStatus(status)
}

// 数据库模型转换为领域模型
func ConvertPaymentDO(thing *po.PaymentPO) *do.PaymentDO {
	if thing == nil {
		return nil
	}

	return &do.PaymentDO{
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
func ConvertPaymentDOs(things []*po.PaymentPO) []*do.PaymentDO {
	if things == nil {
		return nil
	}
	var readers []*do.PaymentDO
	for _, item := range things {
		readers = append(readers, ConvertPaymentDO(item))
	}
	return readers
}
