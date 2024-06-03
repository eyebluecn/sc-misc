package config

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api/subscriptionservice"
)

var (
	SubscriptionRpcClient subscriptionservice.Client
)

func InitRpcClient() {

	subscriptionRpcClient, err := subscriptionservice.NewClient("SubscriptionService", client.WithHostPorts(fmt.Sprintf("0.0.0.0:%v", SubscriptionServerPort)))
	if err != nil {
		klog.CtxInfof(context.Background(), "subscriptionservice client init error: %v", err)
		panic(err)
	}

	SubscriptionRpcClient = subscriptionRpcClient

}
