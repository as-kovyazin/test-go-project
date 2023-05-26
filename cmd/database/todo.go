package database

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
)

type Todo struct {
	ID          int64 `bun:",pk,autoincrement"`
	CreatedAt   int64
	Text        string
	IsCompleted bool
	CompletedAt int64
}

func FindTodoById(id int64, db *bun.DB, ctx context.Context) (todo Todo, err error) {
	err = db.NewSelect().Model(&todo).Where("? = ?", bun.Ident("id"), id).Scan(ctx)
	return
}

func FindUncompletedTodos(db *bun.DB, ctx context.Context) (todos []Todo, err error) {
	if err = db.NewSelect().Model(&todos).Where("? = ?", bun.Ident("is_completed"), false).Scan(ctx); err != nil {
		return
	}
	return
}

func FindCompletedTodos(db *bun.DB, ctx context.Context) (todos []Todo, err error) {
	if err = db.NewSelect().Model(&todos).Where("? = ?", bun.Ident("is_completed"), true).Scan(ctx); err != nil {
		return
	}
	return
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
