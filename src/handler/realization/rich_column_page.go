package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/application"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/do2dto"
	"github.com/eyebluecn/sc-misc/src/converter/dto2do"
	"github.com/eyebluecn/sc-misc/src/model/query"
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
		return errs.CodeErrorf(errs.StatusCodeParamsError, "request 不能为空")
	}
	if request.Operator == nil {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "request.Operator 不能为空")
	}
	if request.Operator.OperatorId == 0 {
		return errs.CodeErrorf(errs.StatusCodeParamsError, "request.Operator.OperatorId 不能为空")
	}

	return nil
}

// 参数校验后的真实处理
func (receiver RichColumnPage) doHandle(ctx context.Context, request sc_misc_api.RichColumnPageRequest) (r *sc_misc_api.RichColumnPageResponse, err error) {

	repoRequest := query.ColumnPageQuery{
		Name:     request.Name,
		AuthorId: request.AuthorId,
		Status:   dto2do.ConvertColumnStatus(request.Status),
		PageNum:  request.PageNum,
		PageSize: request.PageSize,
	}

	richColumns, pagination, err := application.NewColumnReadAppSvc().ReaderColumnList(ctx, *request.Operator, repoRequest)

	return &sc_misc_api.RichColumnPageResponse{
		Data:       do2dto.ConvertRichColumnDTOs(richColumns),
		Pagination: do2dto.ConvertPagination(pagination),
	}, nil

}
