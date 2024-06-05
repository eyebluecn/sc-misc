package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/application"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/do2dto"
)

type PaymentCreate struct{}

func NewPaymentCreate() *PaymentCreate {
	return &PaymentCreate{}
}

// 服务请求入口
func (receiver PaymentCreate) Handle(ctx context.Context, request *sc_misc_api.PaymentCreateRequest) (r *sc_misc_api.PaymentCreateResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver PaymentCreate) CheckParam(ctx context.Context, request *sc_misc_api.PaymentCreateRequest) error {
	if request == nil {
		return errs.BadRequestErrorf("request 不能为空")
	}
	return nil
}

// 参数校验后的真实处理
func (receiver PaymentCreate) doHandle(ctx context.Context, request sc_misc_api.PaymentCreateRequest) (r *sc_misc_api.PaymentCreateResponse, err error) {

	preparePaymentInfo, err := application.NewPaymentWriteApp().Create(ctx, request.OrderNo, request.Method, request.Amount)
	if err != nil {
		return nil, err
	}

	resp := &sc_misc_api.PaymentCreateResponse{
		Data: &sc_misc_api.PaymentPrepareData{
			PaymentDTO:         do2dto.ConvertPaymentDTO(preparePaymentInfo.Payment),
			ThirdTransactionNo: preparePaymentInfo.ThirdTransactionNo,
			NonceStr:           preparePaymentInfo.NonceStr,
		},
		BaseResp: nil,
	}

	return resp, nil
}
