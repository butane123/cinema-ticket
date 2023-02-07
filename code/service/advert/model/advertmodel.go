package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stringx"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AdvertModel = (*customAdvertModel)(nil)
var advertRowsExpectAutoSetButId = strings.Join(stringx.Remove(advertFieldNames, "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`"), ",")

type (
	// AdvertModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdvertModel.
	AdvertModel interface {
		advertModel
		FindAllByIsCom(ctx context.Context, isCom int64) ([]*Advert, error)
		InsertWithNewId(ctx context.Context, data *Advert) (sql.Result, error)
	}

	customAdvertModel struct {
		*defaultAdvertModel
	}
)

// NewAdvertModel returns a model for the database table.
func NewAdvertModel(conn sqlx.SqlConn, c cache.CacheConf) AdvertModel {
	return &customAdvertModel{
		defaultAdvertModel: newAdvertModel(conn, c),
	}
}

func (m *defaultAdvertModel) InsertWithNewId(ctx context.Context, data *Advert) (sql.Result, error) {
	advertIdKey := fmt.Sprintf("%s%v", cacheAdvertIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, advertRowsExpectAutoSetButId)
		return conn.ExecCtx(ctx, query, data.Id, data.Title, data.Content, data.IsCom, data.Status)
	}, advertIdKey)
	return ret, err
}
func (m *defaultAdvertModel) FindAllByIsCom(ctx context.Context, isCom int64) ([]*Advert, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("is_com = ?", isCom).Where("status = ?", 0).ToSql()
	var resp []*Advert
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultAdvertModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(advertRows).From(m.table)
}
