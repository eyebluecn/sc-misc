package mq

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/infrastructure/rpc"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

type Producer struct {
}

func NewProducer() *Producer {
	return &Producer{}
}

// mq发布。一般一个topic会对应一个producer，这里将本系统的发布者一起收到了这里。
// @param tags  消息的tags
// @param keys  消息的Key字段是为了唯一标识消息的，方便运维排查问题。如果不设置Key，则无法定位消息丢失原因。
// @param body  发送的消息体，在智慧课堂中只发送文本消息。
func (mq *Producer) Publish(ctx context.Context, topic MqTopic, tags string, keys string, body string) error {
	klog.CtxInfof(ctx, "模拟发送MQ，topic=%v tags=%v keys=%v body=%v", topic, tags, keys, body)

	//因为模拟的原因，这里直接投递到对应的消费者接口。
	req := &sc_subscription_api.MqMessageArriveRequest{
		Topic: string(topic),
		Tags:  tags,
		Keys:  keys,
		Body:  body,
	}
	err := rpc.NewSubscriptionCaller().MqMessageArrive(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "直接投递MQ消息时出错了: %v", err)
	}

	return nil
}
