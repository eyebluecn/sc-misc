package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/application"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/api_conv"
)

type PaymentQueryById struct{}

func NewPaymentQueryById() *PaymentQueryById {
	return &PaymentQueryById{}
}

// 服务请求入口
func (receiver PaymentQueryById) Handle(ctx context.Context, request *sc_misc_api.PaymentQueryByIdRequest) (r *sc_misc_api.PaymentQueryByIdResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver PaymentQueryById) CheckParam(ctx context.Context, request *sc_misc_api.PaymentQueryByIdRequest) error {
	if request == nil {
		return errs.BadRequestErrorf("request 不能为空")
	}
	if request.PaymentId == 0 {
		return errs.BadRequestErrorf("request 不能为空")
	}
	return nil
}

// 参数校验后的真实处理
func (receiver PaymentQueryById) doHandle(ctx context.Context, request sc_misc_api.PaymentQueryByIdRequest) (r *sc_misc_api.PaymentQueryByIdResponse, err error) {

	payment, err := application.NewPaymentReadApp().QueryById(ctx, request.PaymentId)
	if err != nil {
		return nil, err
	}

	resp := &sc_misc_api.PaymentQueryByIdResponse{
		Data:     api_conv.ConvertPaymentDTO(payment),
		BaseResp: nil,
	}

	return resp, nil
}
