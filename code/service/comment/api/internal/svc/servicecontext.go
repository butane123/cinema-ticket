package svc

import (
	"cinema-ticket/service/comment/api/internal/config"
	"cinema-ticket/service/comment/model"
	"cinema-ticket/service/film/rpc/filmclient"

	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel model.CommentModel
	FilmRpc      filmclient.Film
	RedisClient  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentModel(conn, c.CacheRedis),
		FilmRpc:      filmclient.NewFilm(zrpc.MustNewClient(c.FilmRpc)),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
