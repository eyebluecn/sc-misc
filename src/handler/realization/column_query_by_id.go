package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/application"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/do2dto"
)

type ColumnQueryById struct{}

func NewColumnQueryById() *ColumnQueryById {
	return &ColumnQueryById{}
}

// 服务请求入口
func (receiver ColumnQueryById) Handle(ctx context.Context, request *sc_misc_api.ColumnQueryByIdRequest) (r *sc_misc_api.ColumnQueryByIdResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver ColumnQueryById) CheckParam(ctx context.Context, request *sc_misc_api.ColumnQueryByIdRequest) error {
	if request == nil {
		return errs.BadRequestErrorf("request 不能为空")
	}
	if request.ColumnId == 0 {
		return errs.BadRequestErrorf("request.ColumnId 不能为空")
	}
	return nil
}

// 参数校验后的真实处理
func (receiver ColumnQueryById) doHandle(ctx context.Context, request sc_misc_api.ColumnQueryByIdRequest) (r *sc_misc_api.ColumnQueryByIdResponse, err error) {
	column, err := application.NewColumnReadApp().QueryById(ctx, request.ColumnId)
	if err != nil {
		return nil, err
	}
	response := &sc_misc_api.ColumnQueryByIdResponse{
		Data:     do2dto.ConvertColumnDTO(column),
		BaseResp: nil,
	}

	return response, nil

}
