package api2model_conv

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/model"
)

// 转为枚举
func ConvertColumnStatus(status *sc_misc_api.ColumnStatus) *model.ColumnStatus {
	if status == nil {
		return nil
	}
	result := model.ColumnStatus(*status)
	return &result
}
