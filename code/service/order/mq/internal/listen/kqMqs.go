package listen

import (
	"cinema-ticket/service/order/mq/internal/config"
	kq2 "cinema-ticket/service/order/mq/internal/mqs/kq"
	"cinema-ticket/service/order/mq/internal/svc"
	"context"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

// 创建kafka消息队列
func KqMqs(c config.Config, ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.KqOrderCreate, kq2.NewOrderCreateMq(ctx, svcCtx)),
	}
}
