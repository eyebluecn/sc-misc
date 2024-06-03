package thirdpartypay

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/util"
)

// 第三方支付客户端
type PayClient struct {
}

func NewPayClient() *PayClient {
	return &PayClient{}
}

// 向第三方支付平台下单。
func (mq *PayClient) CreateOrder(ctx context.Context, orderNo string) (*ThirdPartyPayInfo, error) {

	klog.CtxInfof(ctx, "【模拟】向第三方支付平台发起创建订单 orderNo=%v", orderNo)

	info := &ThirdPartyPayInfo{
		OrderNo:            orderNo,
		ThirdTransactionNo: util.RandomString(16),
		NonceStr:           util.RandomString(10),
	}

	klog.CtxInfof(ctx, "【模拟】向第三方支付平台 订单创建成功 orderNo=%v thirdTransactionNo=%v NonceStr=%v", orderNo, info.ThirdTransactionNo, info.NonceStr)

	return info, nil

}

// 向第三方支付平台查询订单。
func (mq *PayClient) QueryOrder(ctx context.Context, orderNo string) (*ThirdPartyPayInfo, error) {

	//随机返回null，或者查到内容。
	if util.RandomBoolean() {

		return nil, nil
	} else {
		klog.CtxInfof(ctx, "【模拟】向第三方支付平台查询订单 orderNo=%v", orderNo)
		info := &ThirdPartyPayInfo{
			OrderNo:            orderNo,
			ThirdTransactionNo: util.RandomString(16),
			NonceStr:           util.RandomString(10),
		}
		klog.CtxInfof(ctx, "【模拟】向第三方支付平台 查询成功 orderNo=%v thirdTransactionNo=%v", orderNo, info.ThirdTransactionNo)
		return info, nil
	}

}
