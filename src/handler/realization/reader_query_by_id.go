package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/application"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/api_conv"
)

type ReaderQueryById struct{}

func NewReaderQueryById() *ReaderQueryById {
	return &ReaderQueryById{}
}

// 服务请求入口
func (receiver ReaderQueryById) Handle(ctx context.Context, request *sc_misc_api.ReaderQueryByIdRequest) (r *sc_misc_api.ReaderQueryByIdResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver ReaderQueryById) CheckParam(ctx context.Context, request *sc_misc_api.ReaderQueryByIdRequest) error {
	if request == nil {
		return errs.BadRequestErrorf("request 不能为空")

	}
	if request.ReaderId == 0 {
		return errs.BadRequestErrorf("request.ReaderId 不能为空")
	}
	return nil
}

// 参数校验后的真实处理
func (receiver ReaderQueryById) doHandle(ctx context.Context, request sc_misc_api.ReaderQueryByIdRequest) (r *sc_misc_api.ReaderQueryByIdResponse, err error) {
	reader, err := application.NewReaderReadApp().QueryById(ctx, request.ReaderId)
	if err != nil {
		return nil, err
	}
	response := &sc_misc_api.ReaderQueryByIdResponse{
		Data:     api_conv.ConvertReaderDTO(reader),
		BaseResp: nil,
	}

	return response, nil
}
