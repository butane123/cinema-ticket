package svc

import (
	"cinema-ticket/service/film/model"
	"cinema-ticket/service/film/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	FilmModel   model.FilmModel
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		FilmModel: model.NewFilmModel(conn, c.CacheRedis),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
