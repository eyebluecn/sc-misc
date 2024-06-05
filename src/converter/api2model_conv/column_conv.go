package api2model_conv

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
)

// 转为枚举
func ConvertColumnStatus(status *sc_misc_api.ColumnStatus) *enums.ColumnStatus {
	if status == nil {
		return nil
	}
	result := enums.ColumnStatus(*status)
	return &result
}
