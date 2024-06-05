package errs

import (
	"fmt"
)

// 项目自定义的错误
type CommonError struct {
	Code StatusCode //错误码
	Msg  string     //错误描述
}

func (receiver CommonError) Error() string {
	return receiver.Msg
}

// 按照格式，产出Err
func Errorf(format string, a ...any) *CommonError {
	msg := fmt.Sprintf(format, a...)
	return &CommonError{Code: StatusCodeUnknown, Msg: msg}
}

// 按照格式，产出带CodeErr
func CodeErrorf(code StatusCode, format string, a ...any) *CommonError {
	msg := fmt.Sprintf(format, a...)
	return &CommonError{Code: code, Msg: msg}
}

// 快速构造特定code的错误
func BadRequestErrorf(format string, a ...any) *CommonError {
	msg := fmt.Sprintf(format, a...)
	return &CommonError{Code: StatusCodeBadRequest, Msg: msg}
}

// 快速构造特定code的错误
func NotFoundErrorf(format string, a ...any) *CommonError {
	msg := fmt.Sprintf(format, a...)
	return &CommonError{Code: StatusCodeNotFound, Msg: msg}
}

// 尝试转为CommonError
func Convert(normalErr error) *CommonError {
	if normalErr == nil {
		return nil
	}

	// 使用类型断言检查错误类型
	if err, ok := normalErr.(*CommonError); ok {
		//是自定义错误就直接返回。
		return err
	} else {
		//不是自定义错误就包装成未知错误返回
		return Errorf(normalErr.Error())
	}

}
