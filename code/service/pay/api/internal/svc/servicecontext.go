package svc

import (
	"cinema-ticket/service/order/rpc/orderclient"
	"cinema-ticket/service/pay/api/internal/config"
	"cinema-ticket/service/pay/model"

	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	PayModel    model.PayModel
	OrderRpc    orderclient.Order
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		PayModel: model.NewPayModel(conn, c.CacheRedis),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
