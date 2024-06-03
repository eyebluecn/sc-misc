package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/application"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/api_conv"
)

type PaymentPaidCallback struct{}

func NewPaymentPaidCallback() *PaymentPaidCallback {
	return &PaymentPaidCallback{}
}

// 服务请求入口
func (receiver PaymentPaidCallback) Handle(ctx context.Context, request *sc_misc_api.PaymentPaidCallbackRequest) (r *sc_misc_api.PaymentPaidCallbackResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver PaymentPaidCallback) CheckParam(ctx context.Context, request *sc_misc_api.PaymentPaidCallbackRequest) error {
	if request == nil {
		return errs.BadRequestErrorf("request 不能为空")
	}
	if request.OrderNo == "" {
		return errs.BadRequestErrorf("request.OrderNo 不能为空")
	}

	return nil
}

// 参数校验后的真实处理
func (receiver PaymentPaidCallback) doHandle(ctx context.Context, request sc_misc_api.PaymentPaidCallbackRequest) (r *sc_misc_api.PaymentPaidCallbackResponse, err error) {

	payment, err := application.NewPaymentWriteApp().PaidCallback(ctx, request.OrderNo)
	if err != nil {
		return nil, err
	}

	resp := &sc_misc_api.PaymentPaidCallbackResponse{
		Data:     api_conv.ConvertPaymentDTO(payment),
		BaseResp: nil,
	}

	return resp, nil

}
