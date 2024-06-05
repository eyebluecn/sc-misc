package application

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/domain"
	"github.com/eyebluecn/sc-misc/src/infrastructure/third_party_service/pay"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/info"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type PaymentWriteApp struct{}

func NewPaymentWriteApp() *PaymentWriteApp {
	return &PaymentWriteApp{}
}

// 创建一个支付单同时返回支付准备物料等信息。
func (receiver PaymentWriteApp) Create(ctx context.Context, orderNo string, method string, amount int64) (*info.PreparePaymentInfo, error) {

	info, err := domain.NewPaymentDomainService().Create(ctx, orderNo, method, amount)
	if err != nil {
		return nil, err
	}

	return info, nil
}

// 准备支付
func (receiver PaymentWriteApp) Prepare(ctx context.Context, paymentId int64) (*info.PreparePaymentInfo, error) {

	payment, err := repo.NewPaymentRepo().CheckById(ctx, paymentId)
	if err != nil {
		return nil, err
	}

	thirdPartyPayInfo, err := pay.NewPayClient().QueryOrder(ctx, payment.OrderNo)
	if err != nil {
		return nil, err
	}
	if thirdPartyPayInfo == nil {
		//说明第三方还没有创建好支付订单。
		thirdPartyPayInfo, err = pay.NewPayClient().CreateOrder(ctx, payment.OrderNo)
		if err != nil {
			return nil, err
		}
	}

	return &info.PreparePaymentInfo{
		Payment:            payment,
		ThirdTransactionNo: thirdPartyPayInfo.ThirdTransactionNo,
		NonceStr:           thirdPartyPayInfo.NonceStr,
	}, nil
}

// 第三方支付平台，支付成功后的回调接口。一班是传输的密文过来，需要用证书解密
func (receiver PaymentWriteApp) PaidCallback(ctx context.Context, orderNo string) (*do.PaymentDO, error) {

	payment, err := repo.NewPaymentRepo().CheckByOrderNo(ctx, orderNo)
	if err != nil {
		return nil, err
	}

	payment, err = domain.NewPaymentDomainService().PaidCallback(ctx, payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
