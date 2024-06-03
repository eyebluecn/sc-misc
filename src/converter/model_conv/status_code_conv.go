package model_conv

import "github.com/eyebluecn/sc-misc/src/common/enums"

// 转为StatusCode
func ConvertStatusCode(statusCode int32) enums.StatusCode {
	return enums.StatusCode(statusCode)
}
