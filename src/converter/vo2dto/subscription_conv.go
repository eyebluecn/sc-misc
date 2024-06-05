package vo2dto

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/model/vo"
	"github.com/eyebluecn/sc-misc/src/model/vo/enums"
)

// 转为枚举
func ConvertSubscriptionStatus(status enums.SubscriptionStatus) sc_misc_api.SubscriptionStatus {
	return sc_misc_api.SubscriptionStatus(status)
}

// 领域模型转为传输模型
func ConvertSubscriptionDTO(thing *vo.SubscriptionVO) *sc_misc_api.SubscriptionDTO {
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
func ConvertSubscriptionDTOs(things []*vo.SubscriptionVO) []*sc_misc_api.SubscriptionDTO {
	if things == nil {
		return nil
	}
	var readerDTOs []*sc_misc_api.SubscriptionDTO
	for _, item := range things {
		readerDTOs = append(readerDTOs, ConvertSubscriptionDTO(item))
	}
	return readerDTOs
}
