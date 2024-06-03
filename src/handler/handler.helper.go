package handler

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_base"
)

// 响应体接口，鸭子类型可用于覆盖所有响应体
type ICommonResponse interface {
	GetBaseResp() *sc_misc_base.BaseResp
	SetBaseResp(*sc_misc_base.BaseResp)
}

// FinalWrapResponse 对错误结果做统一包装处理。 保证最终RPC吐出去的结果没有error.
func FinalWrapResponse(ctx context.Context, resp ICommonResponse, newResp ICommonResponse, err error) ICommonResponse {

	if newResp == nil {
		panic("程序错误！newResp不能为nil")
	}

	if err != nil {
		//如果有错误，那么就使用 newResp
		klog.CtxErrorf(ctx, "遭遇了错误：%v", err)
		newResp.SetBaseResp(&sc_misc_base.BaseResp{StatusCode: 500, StatusMessage: err.Error()})
		return newResp
	} else {

		if resp == nil {
			klog.CtxErrorf(ctx, "返回体不能为空")
			newResp.SetBaseResp(&sc_misc_base.BaseResp{StatusCode: 500, StatusMessage: "返回体异常，不能为空"})
			return newResp
		}

		//如果没有错误，那么就是使用 resp
		if resp.GetBaseResp() == nil {
			//没有设置标准的返回体，将其设置上。
			resp.SetBaseResp(&sc_misc_base.BaseResp{})
		}

		return resp
	}
}
