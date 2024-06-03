package api_conv

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/model/vo_model"
	"github.com/eyebluecn/sc-misc/src/util"
)

// 转为枚举
func ConvertSubscriptionStatus(status vo_model.SubscriptionStatus) sc_misc_api.SubscriptionStatus {
	return sc_misc_api.SubscriptionStatus(status)
}

// 领域模型转为传输模型
func ConvertSubscriptionDTO(thing *vo_model.SubscriptionVO) *sc_misc_api.SubscriptionDTO {
	if thing == nil {
		return nil
	}

	return &sc_misc_api.SubscriptionDTO{
		Id:         thing.ID,
		CreateTime: util.Timestamp(thing.CreateTime),
		UpdateTime: util.Timestamp(thing.UpdateTime),
		ReaderId:   thing.ReaderID,
		ColumnId:   thing.ColumnID,
		OrderId:    thing.OrderID,
		Status:     ConvertSubscriptionStatus(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertSubscriptionDTOs(things []*vo_model.SubscriptionVO) []*sc_misc_api.SubscriptionDTO {
	if things == nil {
		return nil
	}
	var readerDTOs []*sc_misc_api.SubscriptionDTO
	for _, item := range things {
		readerDTOs = append(readerDTOs, ConvertSubscriptionDTO(item))
	}
	return readerDTOs
}
