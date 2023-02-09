package main

import (
	"cinema-ticket/common/errorx"
	"cinema-ticket/service/order/mq/internal/config"
	"cinema-ticket/service/order/mq/internal/listen"
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/service"

	_ "github.com/dtm-labs/driver-gozero"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	// 添加以下自定义错误：
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	//处理所有消息队列
	for _, mq := range listen.Mqs(c) {
		serviceGroup.Add(mq)
	}
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	serviceGroup.Start()
}
