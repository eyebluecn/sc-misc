package handler

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/handler/realization"
)

// 定义一个变量，方便快速查看待实现的方法
var _ sc_misc_api.MiscService

// 定义一个服务实现类
type MiscServiceImpl struct{}

func (receiver *MiscServiceImpl) ColumnOmnibus(ctx context.Context, request *sc_misc_api.ColumnOmnibusRequest) (r *sc_misc_api.ColumnOmnibusResponse, err error) {

	klog.Infof("ColumnOmnibus request: %v", request)
	r, err = realization.NewColumnOmnibus().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_misc_api.ColumnOmnibusResponse{}, err).(*sc_misc_api.ColumnOmnibusResponse), nil
}

func (receiver *MiscServiceImpl) RichColumnPage(ctx context.Context, request *sc_misc_api.RichColumnPageRequest) (r *sc_misc_api.RichColumnPageResponse, err error) {

	klog.Infof("RichColumnPage request: %v", request)
	r, err = realization.NewColumnPage().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_misc_api.RichColumnPageResponse{}, err).(*sc_misc_api.RichColumnPageResponse), nil
}

func (receiver *MiscServiceImpl) ColumnQueryById(ctx context.Context, request *sc_misc_api.ColumnQueryByIdRequest) (r *sc_misc_api.ColumnQueryByIdResponse, err error) {

	klog.Infof("ColumnQueryById request: %v", request)
	r, err = realization.NewColumnQueryById().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_misc_api.ColumnQueryByIdResponse{}, err).(*sc_misc_api.ColumnQueryByIdResponse), nil
}

func (receiver *MiscServiceImpl) ColumnQueryByIds(ctx context.Context, request *sc_misc_api.ColumnQueryByIdsRequest) (r *sc_misc_api.ColumnQueryByIdsResponse, err error) {

	klog.Infof("ColumnQueryByIds request: %v", request)
	r, err = realization.NewColumnQueryByIds().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_misc_api.ColumnQueryByIdsResponse{}, err).(*sc_misc_api.ColumnQueryByIdsResponse), nil
}

func (receiver *MiscServiceImpl) ColumnQuoteQueryByColumnId(ctx context.Context, request *sc_misc_api.ColumnQuoteQueryByColumnIdRequest) (r *sc_misc_api.ColumnQuoteQueryByColumnIdResponse, err error) {

	klog.Infof("ColumnQuoteQueryByColumnId request: %v", request)
	r, err = realization.NewColumnQuoteQueryByColumnId().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_misc_api.ColumnQuoteQueryByColumnIdResponse{}, err).(*sc_misc_api.ColumnQuoteQueryByColumnIdResponse), nil
}

// 登录
func (receiver *MiscServiceImpl) EditorLogin(ctx context.Context, request *sc_misc_api.EditorLoginRequest) (r *sc_misc_api.EditorLoginResponse, err error) {

	klog.Infof("ReaderLogin request: %v", request)
	r, err = realization.NewEditorLogin().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_misc_api.EditorLoginResponse{}, err).(*sc_misc_api.EditorLoginResponse), nil
}

// 登录
func (receiver *MiscServiceImpl) ReaderLogin(ctx context.Context, request *sc_misc_api.ReaderLoginRequest) (r *sc_misc_api.ReaderLoginResponse, err error) {

	klog.Infof("ReaderLogin request: %v", request)
	r, err = realization.NewReaderLogin().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_misc_api.ReaderLoginResponse{}, err).(*sc_misc_api.ReaderLoginResponse), nil
}

func (receiver *MiscServiceImpl) ReaderQueryById(ctx context.Context, request *sc_misc_api.ReaderQueryByIdRequest) (r *sc_misc_api.ReaderQueryByIdResponse, err error) {

	klog.Infof("ReaderQueryById request: %v", request)
	r, err = realization.NewReaderQueryById().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_misc_api.ReaderQueryByIdResponse{}, err).(*sc_misc_api.ReaderQueryByIdResponse), nil
}

func (receiver *MiscServiceImpl) ReaderQueryByIds(ctx context.Context, request *sc_misc_api.ReaderQueryByIdsRequest) (r *sc_misc_api.ReaderQueryByIdsResponse, err error) {
	klog.Infof("ReaderQueryByIds request: %v", request)
	r, err = realization.NewReaderQueryByIds().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_misc_api.ReaderQueryByIdsResponse{}, err).(*sc_misc_api.ReaderQueryByIdsResponse), nil
}

func (receiver *MiscServiceImpl) PaymentQueryById(ctx context.Context, request *sc_misc_api.PaymentQueryByIdRequest) (r *sc_misc_api.PaymentQueryByIdResponse, err error) {
	klog.Infof("PaymentQueryById request: %v", request)
	r, err = realization.NewPaymentQueryById().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_misc_api.PaymentQueryByIdResponse{}, err).(*sc_misc_api.PaymentQueryByIdResponse), nil
}

func (receiver *MiscServiceImpl) PaymentPrepare(ctx context.Context, request *sc_misc_api.PaymentPrepareRequest) (r *sc_misc_api.PaymentPrepareResponse, err error) {
	klog.Infof("PaymentPrepare request: %v", request)
	r, err = realization.NewPaymentPrepare().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_misc_api.PaymentPrepareResponse{}, err).(*sc_misc_api.PaymentPrepareResponse), nil
}
func (receiver *MiscServiceImpl) PaymentCreate(ctx context.Context, request *sc_misc_api.PaymentCreateRequest) (r *sc_misc_api.PaymentCreateResponse, err error) {
	klog.Infof("PaymentCreate request: %v", request)
	r, err = realization.NewPaymentCreate().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_misc_api.PaymentCreateResponse{}, err).(*sc_misc_api.PaymentCreateResponse), nil
}
func (receiver *MiscServiceImpl) PaymentPaidCallback(ctx context.Context, request *sc_misc_api.PaymentPaidCallbackRequest) (r *sc_misc_api.PaymentPaidCallbackResponse, err error) {
	klog.Infof("PaymentPaidCallback request: %v", request)
	r, err = realization.NewPaymentPaidCallback().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_misc_api.PaymentPaidCallbackResponse{}, err).(*sc_misc_api.PaymentPaidCallbackResponse), nil
}

func (receiver *MiscServiceImpl) PaymentPublishMq(ctx context.Context, request *sc_misc_api.PaymentPublishMqRequest) (r *sc_misc_api.PaymentPublishMqResponse, err error) {
	klog.Infof("PaymentPublishMq request: %v", request)
	r, err = realization.NewPaymentPublishMq().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_misc_api.PaymentPublishMqResponse{}, err).(*sc_misc_api.PaymentPublishMqResponse), nil
}
