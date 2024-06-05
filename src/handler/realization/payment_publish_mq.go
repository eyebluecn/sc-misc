package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/errs"
)

type PaymentPublishMq struct{}

func NewPaymentPublishMq() *PaymentPublishMq {
	return &PaymentPublishMq{}
}

// 服务请求入口
func (receiver PaymentPublishMq) Handle(ctx context.Context, request *sc_misc_api.PaymentPublishMqRequest) (r *sc_misc_api.PaymentPublishMqResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver PaymentPublishMq) CheckParam(ctx context.Context, request *sc_misc_api.PaymentPublishMqRequest) error {
	if request == nil {
		return errs.BadRequestErrorf("request 不能为空")
	}

	return nil
}

// 参数校验后的真实处理
func (receiver PaymentPublishMq) doHandle(ctx context.Context, request sc_misc_api.PaymentPublishMqRequest) (r *sc_misc_api.PaymentPublishMqResponse, err error) {

	return nil, errs.CodeErrorf(errs.StatusCodeServer, "不支持直接调用RPC接口发送MQ消息")

}
