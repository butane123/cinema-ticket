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

var _ PayModel = (*customPayModel)(nil)

var payRowsExpectAutoSetButId = strings.Join(stringx.Remove(payFieldNames, "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), ",")

type (
	// PayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPayModel.
	PayModel interface {
		payModel
		FindByOid(ctx context.Context, orderId int64) (*Pay, error)
		InsertWithNewId(ctx context.Context, data *Pay) (sql.Result, error)
	}

	customPayModel struct {
		*defaultPayModel
	}
)

func (m *defaultPayModel) InsertWithNewId(ctx context.Context, data *Pay) (sql.Result, error) {
	payIdKey := fmt.Sprintf("%s%v", cachePayIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, payRowsExpectAutoSetButId)
		return conn.ExecCtx(ctx, query, data.Id, data.Uid, data.Oid, data.Amount, data.Source, data.Status)
	}, payIdKey)
	return ret, err
}

// NewPayModel returns a model for the database table.
func NewPayModel(conn sqlx.SqlConn, c cache.CacheConf) PayModel {
	return &customPayModel{
		defaultPayModel: newPayModel(conn, c),
	}
}

func (m *defaultPayModel) FindByOid(ctx context.Context, orderId int64) (*Pay, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("oid = ?", orderId).ToSql()
	var resp Pay
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultPayModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(payRows).From(m.table)
}
