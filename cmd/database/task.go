package database

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
)

type Task struct {
	ID          int64 `bun:",pk,autoincrement"`
	CreatedAt   int64
	Text        string
	IsCompleted bool
	CompletedAt int64
}

func FindTaskById(id int64, db *bun.DB, ctx context.Context) (task Task, err error) {
	err = db.NewSelect().Model(&task).Where("? = ?", bun.Ident("id"), id).Scan(ctx)
	return
}

func FindUncompletedTasks(db *bun.DB, ctx context.Context) (tasks []Task, err error) {
	if err = db.NewSelect().Model(&tasks).Where("? = ?", bun.Ident("is_completed"), false).Scan(ctx); err != nil {
		return
	}
	return
}

func FindCompletedTasks(db *bun.DB, ctx context.Context) (tasks []Task, err error) {
	if err = db.NewSelect().Model(&tasks).Where("? = ?", bun.Ident("is_completed"), true).Scan(ctx); err != nil {
		return
	}
	return
}

func (task *Task) Insert(db *bun.DB, ctx context.Context) (sql.Result, error) {
	return db.NewInsert().Model(task).Exec(ctx)
}

func (task *Task) Update(db *bun.DB, ctx context.Context) (sql.Result, error) {
	return db.NewUpdate().Model(task).WherePK().Exec(ctx)
}

func (task *Task) Delete(db *bun.DB, ctx context.Context) (sql.Result, error) {
	return db.NewDelete().Model(task).WherePK().Exec(ctx)
}
