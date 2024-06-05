package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/application"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/do2dto"
)

type ColumnOmnibus struct{}

func NewColumnOmnibus() *ColumnOmnibus {
	return &ColumnOmnibus{}
}

// 服务请求入口
func (receiver ColumnOmnibus) Handle(ctx context.Context, request *sc_misc_api.ColumnOmnibusRequest) (r *sc_misc_api.ColumnOmnibusResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver ColumnOmnibus) CheckParam(ctx context.Context, request *sc_misc_api.ColumnOmnibusRequest) error {
	if request == nil {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "request 不能为空")
	}
	if request.AuthorName == "" {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "request.AuthorName 不能为空")
	}
	if request.ColumnName == "" {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "request.ColumnName 不能为空")
	}
	if request.ColumnPrice <= 0 {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "request.ColumnPrice 必须是正数")
	}

	if request.Operator == nil {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "request.Operator 必填")
	}

	return nil
}

// 参数校验后的真实处理
func (receiver ColumnOmnibus) doHandle(ctx context.Context, request sc_misc_api.ColumnOmnibusRequest) (r *sc_misc_api.ColumnOmnibusResponse, err error) {

	richColumn, err := application.NewColumnWriteApp().ColumnOmnibus(ctx, request)
	if err != nil {
		return nil, err
	}

	return &sc_misc_api.ColumnOmnibusResponse{
		RichColumnDTO: do2dto.ConvertRichColumnDTO(richColumn),
	}, nil
}
