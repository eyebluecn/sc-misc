package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/application"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/do2dto"
)

type PaymentPrepare struct{}

func NewPaymentPrepare() *PaymentPrepare {
	return &PaymentPrepare{}
}

// 服务请求入口
func (receiver PaymentPrepare) Handle(ctx context.Context, request *sc_misc_api.PaymentPrepareRequest) (r *sc_misc_api.PaymentPrepareResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver PaymentPrepare) CheckParam(ctx context.Context, request *sc_misc_api.PaymentPrepareRequest) error {
	if request == nil {
		return errs.BadRequestErrorf("request 不能为空")
	}
	if request.PaymentId == 0 {
		return errs.BadRequestErrorf("request.PaymentId 不能为空")
	}

	return nil
}

// 参数校验后的真实处理
func (receiver PaymentPrepare) doHandle(ctx context.Context, request sc_misc_api.PaymentPrepareRequest) (r *sc_misc_api.PaymentPrepareResponse, err error) {

	preparePaymentInfo, err := application.NewPaymentWriteApp().Prepare(ctx, request.PaymentId)
	if err != nil {
		return nil, err
	}

	resp := &sc_misc_api.PaymentPrepareResponse{
		Data: &sc_misc_api.PaymentPrepareData{
			PaymentDTO:         do2dto.ConvertPaymentDTO(preparePaymentInfo.Payment),
			ThirdTransactionNo: preparePaymentInfo.ThirdTransactionNo,
			NonceStr:           preparePaymentInfo.NonceStr,
		},
		BaseResp: nil,
	}

	return resp, nil
}
