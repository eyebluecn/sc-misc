package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/application"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/do2dto"
)

type ReaderQueryByIds struct{}

func NewReaderQueryByIds() *ReaderQueryByIds {
	return &ReaderQueryByIds{}
}

// 服务请求入口
func (receiver ReaderQueryByIds) Handle(ctx context.Context, request *sc_misc_api.ReaderQueryByIdsRequest) (r *sc_misc_api.ReaderQueryByIdsResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver ReaderQueryByIds) CheckParam(ctx context.Context, request *sc_misc_api.ReaderQueryByIdsRequest) error {
	if request == nil {
		return errs.BadRequestErrorf("request 不能为空")

	}
	if len(request.ReaderIds) == 0 {
		return errs.BadRequestErrorf("request.ReaderId 不能为空")
	}
	return nil
}

// 参数校验后的真实处理
func (receiver ReaderQueryByIds) doHandle(ctx context.Context, request sc_misc_api.ReaderQueryByIdsRequest) (r *sc_misc_api.ReaderQueryByIdsResponse, err error) {
	readers, err := application.NewReaderReadApp().QueryByIds(ctx, request.ReaderIds)
	if err != nil {
		return nil, err
	}
	response := &sc_misc_api.ReaderQueryByIdsResponse{
		Data:     do2dto.ConvertReaderDTOs(readers),
		BaseResp: nil,
	}

	return response, nil
}
