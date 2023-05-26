package database

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
)

type Todo struct {
	ID          int64 `bun:",pk,autoincrement"`
	Text        string
	IsCompleted bool
}

func (todo *Todo) Insert(db *bun.DB, ctx context.Context) (sql.Result, error) {
	return db.NewInsert().Model(todo).Exec(ctx)
}

func (todo *Todo) Update(db *bun.DB, ctx context.Context) (sql.Result, error) {
	return db.NewUpdate().Model(todo).WherePK().Exec(ctx)
}

func (todo *Todo) Delete(db *bun.DB, ctx context.Context) (sql.Result, error) {
	return db.NewDelete().Model(todo).WherePK().Exec(ctx)
}
