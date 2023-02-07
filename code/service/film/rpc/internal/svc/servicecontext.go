package svc

import (
	"cinema-ticket/service/film/model"
	"cinema-ticket/service/film/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	FilmModel model.FilmModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		FilmModel: model.NewFilmModel(conn, c.CacheRedis),
	}
}
