package do2po

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 枚举转为存储的整数
func ConvertPaymentStatus(status enums.PaymentStatus) int32 {
	return int32(status)
}

func ConvertPaymentPO(thing *do.PaymentDO) *po.PaymentPO {
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
		Status:             ConvertPaymentStatus(thing.Status),
	}
}

func ConvertPaymentPOs(things []*do.PaymentDO) []*po.PaymentPO {
	if things == nil {
		return nil
	}
	var results []*po.PaymentPO
	for _, item := range things {
		results = append(results, ConvertPaymentPO(item))
	}
	return results
}
