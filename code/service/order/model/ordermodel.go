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

var _ OrderModel = (*customOrderModel)(nil)
var orderRowsExpectAutoSetButId = strings.Join(stringx.Remove(orderFieldNames, "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), ",")

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		FindAllByUid(ctx context.Context, userId int64) ([]*Order, error)
		FindAllPaidByUid(ctx context.Context, userId int64) ([]*Order, error)
		FindLatestByUid(ctx context.Context, userId int64) (*Order, error)
		TxInsert(tx *sql.Tx, data *Order) (sql.Result, error)
		TxDelete(tx *sql.Tx, id int64) error
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn, c cache.CacheConf) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn, c),
	}
}

func (m *defaultOrderModel) FindAllByUid(ctx context.Context, userId int64) ([]*Order, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("uid = ?", userId).ToSql()
	var resp []*Order
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

func (m *defaultOrderModel) FindAllPaidByUid(ctx context.Context, userId int64) ([]*Order, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("uid = ?", userId).Where("status = 1").ToSql()
	var resp []*Order
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
func (m *defaultOrderModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(orderRows).From(m.table)
}

func (m *defaultOrderModel) TxInsert(tx *sql.Tx, data *Order) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, orderRowsExpectAutoSetButId)
	ret, err := tx.Exec(query, data.Id, data.Uid, data.Fid, data.Amount, data.Status)
	return ret, err
}

func (m *defaultOrderModel) TxDelete(tx *sql.Tx, id int64) error {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return tx.Exec(query, id)
	}, orderIdKey)
	return err
}
func (m *defaultOrderModel) FindLatestByUid(ctx context.Context, userId int64) (*Order, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("uid = ?", userId).OrderBy("create_time desc").Limit(1).ToSql()
	var resp Order
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
