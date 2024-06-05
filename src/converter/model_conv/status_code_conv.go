package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/common/errs"
)

// 转为StatusCode
func ConvertStatusCode(statusCode int32) errs.StatusCode {
	return errs.StatusCode(statusCode)
}
