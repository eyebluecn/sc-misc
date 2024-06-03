package main

import (
	"fmt"
	"github.com/cloudwego/kitex/server"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api/miscservice"
	"github.com/eyebluecn/sc-misc/src/common/config"
	"github.com/eyebluecn/sc-misc/src/handler"
	"log"
	"net"
	"time"
)

func main() {

	//初始化MySQL
	config.InitMySQL()

	//自定义端口
	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%v", config.MiscServerPort))
	svr := miscservice.NewServer(new(handler.MiscServiceImpl),
		server.WithServiceAddr(addr),                  //地址端口
		server.WithReadWriteTimeout(365*24*time.Hour), //一段时间没有读到客户端的请求，就报超时错误，设置1年。 https://developer.aliyun.com/article/1408192
		server.WithExitWaitTime(5*time.Second),        //Server 在收到退出信号时的等待时间 可理解为优雅关闭连接的时长。
	)

	//初始化RPC客户端
	config.InitRpcClient()

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
