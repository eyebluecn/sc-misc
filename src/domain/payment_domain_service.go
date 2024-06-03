package domain

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/api_conv"
	"github.com/eyebluecn/sc-misc/src/infra/mq"
	"github.com/eyebluecn/sc-misc/src/infra/thirdpartypay"
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
	"github.com/eyebluecn/sc-misc/src/util"
	"time"
)

type PaymentDomainService struct{}

func NewPaymentDomainService() *PaymentDomainService {
	return &PaymentDomainService{}
}

// 创建支付单，同时做支付的准备。
func (receiver PaymentDomainService) Create(ctx context.Context, orderNo string, method string, amount int64) (*model.PreparePaymentInfo, error) {
	payment, err := repo.NewPaymentRepo().QueryByOrderNo(ctx, orderNo)
	if err != nil {
		return nil, err
	}
	if payment != nil {
		return nil, errs.BadRequestErrorf("订单号 %v 对应的支付单已存在，请勿重复创建", orderNo)
	}

	//重新构建新的支付单领域对象。
	payment = &model.Payment{
		ID:                 0,
		CreateTime:         time.Now(),
		UpdateTime:         time.Now(),
		OrderNo:            orderNo,
		Method:             method,
		ThirdTransactionNo: "",
		Amount:             amount,
		Status:             model.PaymentStatusUnpaid,
	}
	//插入数据库
	payment, err = repo.NewPaymentRepo().Insert(ctx, payment)
	if err != nil {
		return nil, err
	}

	//查询第三方支付信息。
	thirdPartyPayInfo, err := thirdpartypay.NewPayClient().QueryOrder(ctx, orderNo)
	if err != nil {
		return nil, err
	}
	if thirdPartyPayInfo == nil {
		//说明还没有去第三方支付平台下单。
		thirdPartyPayInfo, err = thirdpartypay.NewPayClient().CreateOrder(ctx, orderNo)
		if err != nil {
			return nil, err
		}
		if thirdPartyPayInfo == nil {
			return nil, errs.NotFoundErrorf("三方支付平台无法创建订单%v", orderNo)
		}
	}

	//准备返回结果.
	info := &model.PreparePaymentInfo{
		Payment:            payment,
		ThirdTransactionNo: thirdPartyPayInfo.ThirdTransactionNo,
		NonceStr:           thirdPartyPayInfo.NonceStr,
	}

	return info, nil
}

// 支付成功的回调函数。
func (receiver PaymentDomainService) PaidCallback(ctx context.Context, payment *model.Payment) (*model.Payment, error) {

	//修改支付单状态
	affectedRows, err := repo.NewPaymentRepo().UpdateStatus(ctx, payment.ID, model.PaymentStatusPaid)
	if err != nil {
		return nil, err
	}
	klog.CtxInfof(ctx, "更新paymentId=%v的状态成功，affectedRows=%v", payment.ID, affectedRows)

	//查询一次最新状态
	payment, err = repo.NewPaymentRepo().CheckByOrderNo(ctx, payment.OrderNo)
	if err != nil {
		return nil, err
	}

	//发送领域事件
	err = receiver.SendMqPaid(ctx, payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

// 发送支付成功的领域事件
func (receiver PaymentDomainService) SendMqPaid(ctx context.Context, payment *model.Payment) error {
	//发布领域事件。
	traceId := util.Uuid()
	event := sc_misc_api.PaymentMqEvent_PAYMENT_PAID
	payload := &sc_misc_api.PaymentMqPayload{
		Topic:      string(mq.MqTopicPayment),
		Tags:       event.String(),
		Keys:       traceId,
		Event:      event,
		OccurTime:  util.Timestamp(time.Now()),
		PaymentDTO: api_conv.ConvertPaymentDTO(payment),
	}
	err := mq.NewProducer().Publish(ctx, mq.MqTopicPayment, event.String(), traceId, util.ToJSON(payload))
	if err != nil {
		return err
	}

	return nil
}
