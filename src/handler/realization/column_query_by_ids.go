package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/application"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/do2dto"
)

type ColumnQueryByIds struct{}

func NewColumnQueryByIds() *ColumnQueryByIds {
	return &ColumnQueryByIds{}
}

// 服务请求入口
func (receiver ColumnQueryByIds) Handle(ctx context.Context, request *sc_misc_api.ColumnQueryByIdsRequest) (r *sc_misc_api.ColumnQueryByIdsResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver ColumnQueryByIds) CheckParam(ctx context.Context, request *sc_misc_api.ColumnQueryByIdsRequest) error {
	if request == nil {
		return errs.BadRequestErrorf("request 不能为空")
	}
	if len(request.ColumnIds) == 0 {
		return errs.BadRequestErrorf("request.ColumnIds 不能为空")
	}
	return nil
}

// 参数校验后的真实处理
func (receiver ColumnQueryByIds) doHandle(ctx context.Context, request sc_misc_api.ColumnQueryByIdsRequest) (r *sc_misc_api.ColumnQueryByIdsResponse, err error) {
	columns, err := application.NewColumnReadAppSvc().QueryByIds(ctx, request.ColumnIds)
	if err != nil {
		return nil, err
	}
	response := &sc_misc_api.ColumnQueryByIdsResponse{
		Data:     do2dto.ConvertColumnDTOs(columns),
		BaseResp: nil,
	}

	return response, nil

}
