package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/application"
	"github.com/eyebluecn/sc-misc/src/common/enums"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/api2model_conv"
	"github.com/eyebluecn/sc-misc/src/converter/api_conv"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type RichColumnPage struct{}

func NewColumnPage() *RichColumnPage {
	return &RichColumnPage{}
}

// 服务请求入口
func (receiver RichColumnPage) Handle(ctx context.Context, request *sc_misc_api.RichColumnPageRequest) (r *sc_misc_api.RichColumnPageResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver RichColumnPage) CheckParam(ctx context.Context, request *sc_misc_api.RichColumnPageRequest) error {
	if request == nil {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "request 不能为空")
	}
	if request.Operator == nil {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "request.Operator 不能为空")
	}
	if request.Operator.OperatorId == 0 {
		return errs.CodeErrorf(enums.StatusCodeParamsError, "request.Operator.OperatorId 不能为空")
	}

	return nil
}

// 参数校验后的真实处理
func (receiver RichColumnPage) doHandle(ctx context.Context, request sc_misc_api.RichColumnPageRequest) (r *sc_misc_api.RichColumnPageResponse, err error) {

	repoRequest := repo.ColumnPageRequest{
		Name:     request.Name,
		AuthorId: request.AuthorId,
		Status:   api2model_conv.ConvertColumnStatus(request.Status),
		PageNum:  request.PageNum,
		PageSize: request.PageSize,
	}

	richColumns, pagination, err := application.NewColumnReadApp().ReaderColumnList(ctx, *request.Operator, repoRequest)

	return &sc_misc_api.RichColumnPageResponse{
		Data:       api_conv.ConvertRichColumnDTOs(richColumns),
		Pagination: api_conv.ConvertPagination(pagination),
	}, nil

}
