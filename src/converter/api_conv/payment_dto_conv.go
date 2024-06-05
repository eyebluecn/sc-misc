package api_conv

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
)

// 转为枚举
func ConvertPaymentStatus(status enums.PaymentStatus) sc_misc_api.PaymentStatus {
	return sc_misc_api.PaymentStatus(status)
}

// 领域模型转为传输模型
func ConvertPaymentDTO(thing *do.Payment) *sc_misc_api.PaymentDTO {
	if thing == nil {
		return nil
	}

	return &sc_misc_api.PaymentDTO{
		Id:                 thing.ID,
		CreateTime:         util.Timestamp(thing.CreateTime),
		UpdateTime:         util.Timestamp(thing.UpdateTime),
		OrderNo:            thing.OrderNo,
		Method:             thing.Method,
		ThirdTransactionNo: thing.ThirdTransactionNo,
		Amount:             thing.Amount,
		Status:             ConvertPaymentStatus(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertPaymentDTOs(things []*do.Payment) []*sc_misc_api.PaymentDTO {
	if things == nil {
		return nil
	}
	var resultList []*sc_misc_api.PaymentDTO
	for _, item := range things {
		resultList = append(resultList, ConvertPaymentDTO(item))
	}
	return resultList
}
