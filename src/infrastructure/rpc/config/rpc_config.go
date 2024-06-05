package config

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/config"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api/subscriptionservice"
	"sync"
)

var (
	SubscriptionRpcClient subscriptionservice.Client
)

type RpcConfig struct {
}

var (
	rpcConfig     *RpcConfig
	rpcConfigOnce sync.Once
)

// 采用单例模型
func DefaultMysqlConfig() *RpcConfig {
	rpcConfigOnce.Do(func() {
		rpcConfig = &RpcConfig{}
	})
	return rpcConfig
}

// 初始化客户端
func (receiver *RpcConfig) Init() {

	subscriptionRpcClient, err := subscriptionservice.NewClient("SubscriptionService", client.WithHostPorts(fmt.Sprintf("0.0.0.0:%v", config.SubscriptionServerPort)))
	if err != nil {
		klog.CtxInfof(context.Background(), "subscriptionservice client init error: %v", err)
		panic(err)
	}

	SubscriptionRpcClient = subscriptionRpcClient

}
