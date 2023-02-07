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

var _ CommentModel = (*customCommentModel)(nil)
var commentRowsExpectAutoSetButId = strings.Join(stringx.Remove(commentFieldNames, "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		FindAllByUid(ctx context.Context, userId int64) ([]*Comment, error)
		InsertWithNewId(ctx context.Context, data *Comment) (sql.Result, error)
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn sqlx.SqlConn, c cache.CacheConf) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn, c),
	}
}

func (m *defaultCommentModel) InsertWithNewId(ctx context.Context, data *Comment) (sql.Result, error) {
	commentIdKey := fmt.Sprintf("%s%v", cacheCommentIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, commentRowsExpectAutoSetButId)
		return conn.ExecCtx(ctx, query, data.Id, data.Uid, data.Fid, data.Title, data.Content, data.IsAnonymous)
	}, commentIdKey)
	return ret, err
}

func (m *defaultCommentModel) FindAllByUid(ctx context.Context, userId int64) ([]*Comment, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("uid = ?", userId).ToSql()
	var resp []*Comment
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
func (m *defaultCommentModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(commentRows).From(m.table)
}
