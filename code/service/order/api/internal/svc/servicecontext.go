package svc

import (
	"cinema-ticket/service/film/rpc/filmclient"
	"cinema-ticket/service/order/api/internal/config"
	"cinema-ticket/service/order/model"
	"cinema-ticket/service/order/rpc/orderclient"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	RedisClient         *redis.Redis
	OrderModel          model.OrderModel
	FilmRpc             filmclient.Film
	OrderRpc            orderclient.Order
	KqOrderCreateClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		OrderModel:          model.NewOrderModel(conn, c.CacheRedis),
		FilmRpc:             filmclient.NewFilm(zrpc.MustNewClient(c.FilmRpc)),
		OrderRpc:            orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		KqOrderCreateClient: kq.NewPusher(c.KqOrderCreate.Brokers, c.KqOrderCreate.Topic),
	}
}
