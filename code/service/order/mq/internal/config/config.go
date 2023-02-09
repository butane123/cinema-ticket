package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Mysql struct {
		DataSource string
	}
	CacheRedis    cache.CacheConf
	Redis         redis.RedisConf
	FilmRpc       zrpc.RpcClientConf
	OrderRpc      zrpc.RpcClientConf
	DtmServer     string
	KqOrderCreate kq.KqConf
}
