package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/dto2result"
	"github.com/eyebluecn/sc-misc/src/converter/dto2vo"
	"github.com/eyebluecn/sc-misc/src/converter/po2do"
	"github.com/eyebluecn/sc-misc/src/infrastructure/rpc/config"
	"github.com/eyebluecn/sc-misc/src/model/result"
	"github.com/eyebluecn/sc-misc/src/model/vo"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

type SubscriptionCaller struct {
}

func NewSubscriptionCaller() SubscriptionCaller {
	return SubscriptionCaller{}
}

// 获取订阅分页列表。预期获得获取订阅分页列表，其余一律算错误。
// 如果err==nil，则ReaderVO!=nil
func (receiver SubscriptionCaller) SubscriptionPage(ctx context.Context, request *sc_subscription_api.SubscriptionPageRequest) ([]*vo.SubscriptionVO, *result.Pagination, error) {
	response, err := config.SubscriptionRpcClient.SubscriptionPage(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "Reader login failed: %v", err)
		return nil, nil, err
	}
	if response == nil {
		klog.CtxErrorf(ctx, "Failed: response is nil")
		return nil, nil, errs.Errorf("response is nil")
	}
	if response.BaseResp == nil {
		klog.CtxErrorf(ctx, "Failed: response.BaseResp is nil")
		return nil, nil, errs.Errorf("response is nil")
	}
	statusCode := po2do.ConvertStatusCode(response.BaseResp.StatusCode)
	if statusCode != errs.StatusCodeOk {
		return nil, nil, errs.CodeErrorf(statusCode, response.BaseResp.StatusMessage)
	}
	if response.Data == nil {
		return nil, nil, errs.CodeErrorf(errs.StatusCodeNotFound, "Failed: data is nil.")
	}
	if response.Pagination == nil {
		return nil, nil, errs.CodeErrorf(errs.StatusCodeNotFound, "Failed: pagination is nil.")
	}

	readerVos := dto2vo.ConvertSubscriptionVOs(response.Data)
	pagination := dto2result.ConvertSubscriptionPagination(response.Pagination)

	return readerVos, pagination, nil
}

// 获取某个读者对于一个专栏的订阅情况
func (receiver SubscriptionCaller) SubscriptionList(ctx context.Context, readerId int64, columnIds []int64) ([]*vo.SubscriptionVO, error) {

	request := &sc_subscription_api.SubscriptionPageRequest{
		PageNum:   1,
		PageSize:  100,
		ReaderId:  &readerId,
		ColumnIds: columnIds,
	}

	page, _, err := receiver.SubscriptionPage(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "SubscriptionPage failed: %v", err)
		return nil, err
	}

	return page, nil
}

// 直接投递Mq消息过去
func (receiver SubscriptionCaller) MqMessageArrive(ctx context.Context, request *sc_subscription_api.MqMessageArriveRequest) error {
	response, err := config.SubscriptionRpcClient.MqMessageArrive(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "MqMessageArrive failed: %v", err)
		return err
	}

	if response == nil {
		klog.CtxErrorf(ctx, "Failed: response is nil")
		return errs.Errorf("response is nil")
	}
	if response.BaseResp == nil {
		klog.CtxErrorf(ctx, "Failed: response.BaseResp is nil")
		return errs.Errorf("response is nil")
	}
	statusCode := po2do.ConvertStatusCode(response.BaseResp.StatusCode)
	if statusCode != errs.StatusCodeOk {
		return errs.CodeErrorf(statusCode, response.BaseResp.StatusMessage)
	}

	return nil
}
