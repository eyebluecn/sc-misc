package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/api_conv"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type ColumnQuoteQueryByColumnId struct{}

func NewColumnQuoteQueryByColumnId() *ColumnQuoteQueryByColumnId {
	return &ColumnQuoteQueryByColumnId{}
}

// 服务请求入口
func (receiver ColumnQuoteQueryByColumnId) Handle(ctx context.Context, request *sc_misc_api.ColumnQuoteQueryByColumnIdRequest) (r *sc_misc_api.ColumnQuoteQueryByColumnIdResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver ColumnQuoteQueryByColumnId) CheckParam(ctx context.Context, request *sc_misc_api.ColumnQuoteQueryByColumnIdRequest) error {
	if request == nil {
		return errs.BadRequestErrorf("request 不能为空")
	}
	if request.ColumnId == 0 {
		return errs.BadRequestErrorf("request.ColumnId 不能为空")
	}
	return nil
}

// 参数校验后的真实处理
func (receiver ColumnQuoteQueryByColumnId) doHandle(ctx context.Context, request sc_misc_api.ColumnQuoteQueryByColumnIdRequest) (r *sc_misc_api.ColumnQuoteQueryByColumnIdResponse, err error) {
	columnQuote, err := repo.NewColumnQuoteRepo().QueryByColumnId(ctx, request.ColumnId)
	if err != nil {
		return nil, err
	}
	response := &sc_misc_api.ColumnQuoteQueryByColumnIdResponse{
		Data:     api_conv.ConvertColumnQuoteDTO(columnQuote),
		BaseResp: nil,
	}

	return response, nil

}
