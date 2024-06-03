package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/application"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/api_conv"
)

type EditorLogin struct{}

func NewEditorLogin() *EditorLogin {
	return &EditorLogin{}
}

// 服务请求入口
func (receiver EditorLogin) Handle(ctx context.Context, request *sc_misc_api.EditorLoginRequest) (r *sc_misc_api.EditorLoginResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver EditorLogin) CheckParam(ctx context.Context, request *sc_misc_api.EditorLoginRequest) error {
	if request == nil {
		return errs.BadRequestErrorf("request 不能为空")
	}
	return nil
}

// 参数校验后的真实处理
func (receiver EditorLogin) doHandle(ctx context.Context, request sc_misc_api.EditorLoginRequest) (r *sc_misc_api.EditorLoginResponse, err error) {

	editor, err := application.NewEditorReadApp().Login(ctx, request.Username, request.Password)
	if err != nil {
		klog.CtxErrorf(ctx, "登录过程出错：%v", err)
		return nil, err
	}

	return &sc_misc_api.EditorLoginResponse{
		Data: api_conv.ConvertEditorDTO(editor),
	}, nil
}
