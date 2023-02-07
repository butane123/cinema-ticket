package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stringx"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FilmModel = (*customFilmModel)(nil)
var filmRowsExpectAutoSetButId = strings.Join(stringx.Remove(filmFieldNames, "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), ",")

type (
	// FilmModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFilmModel.
	FilmModel interface {
		filmModel
		TxUpdate(tx *sql.Tx, data *Film) error
		InsertWithNewId(ctx context.Context, data *Film) (sql.Result, error)
	}

	customFilmModel struct {
		*defaultFilmModel
	}
)

// NewFilmModel returns a model for the database table.
func NewFilmModel(conn sqlx.SqlConn, c cache.CacheConf) FilmModel {
	return &customFilmModel{
		defaultFilmModel: newFilmModel(conn, c),
	}
}

func (m *defaultFilmModel) InsertWithNewId(ctx context.Context, data *Film) (sql.Result, error) {
	filmIdKey := fmt.Sprintf("%s%v", cacheFilmIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, filmRowsExpectAutoSetButId)
		return conn.ExecCtx(ctx, query, data.Id, data.Name, data.Desc, data.Stock, data.Amount, data.Screenwriter, data.Director, data.Length, data.IsSelectSeat)
	}, filmIdKey)
	return ret, err
}

func (m *defaultFilmModel) TxUpdate(tx *sql.Tx, data *Film) error {
	filmIdKey := fmt.Sprintf("%s%v", cacheFilmIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, filmRowsWithPlaceHolder)
		return tx.Exec(query, data.Name, data.Desc, data.Stock, data.Amount, data.Screenwriter, data.Director, data.Length, data.IsSelectSeat, data.Id)
	}, filmIdKey)
	return err
}
