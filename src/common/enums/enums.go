package enums

type StatusCode int32

const (
	StatusCodeOk           StatusCode = 0    //成功
	StatusCodeLogin        StatusCode = 4001 //没有登录
	StatusCodeNotFound     StatusCode = 4002 //资源没找到
	StatusCodeBadRequest   StatusCode = 4003 //请求不合法
	StatusCodeUnauthorized StatusCode = 4004 //没有权限
	StatusCodeParamsError  StatusCode = 4005 //参数错误
	StatusCodeServer       StatusCode = 5000 //服务器内部出错
	StatusCodeUnknown      StatusCode = 5001 //服务器未知错误
)
