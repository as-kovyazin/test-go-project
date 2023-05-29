package repositories

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"iSpringTest/database"
)

type Task struct {
	DBConnection *bun.DB
}

func CreateTaskRepository(db *bun.DB) *Task {
	return &Task{
		DBConnection: db,
	}
}

func (tr Task) FindTaskById(id int64, ctx context.Context) (task database.Task, err error) {
	err = tr.DBConnection.NewSelect().Model(&task).Where("? = ?", bun.Ident("id"), id).Scan(ctx)
	return
}

func (tr Task) FindUncompletedTasks(ctx context.Context) (tasks []database.Task, err error) {
	if err = tr.DBConnection.NewSelect().Model(&tasks).Where("? = ?", bun.Ident("is_completed"), false).Scan(ctx); err != nil {
		return
	}
	return
}

func (tr Task) FindCompletedTasks(ctx context.Context) (tasks []database.Task, err error) {
	if err = tr.DBConnection.NewSelect().Model(&tasks).Where("? = ?", bun.Ident("is_completed"), true).Scan(ctx); err != nil {
		return
	}
	return
}

func (tr Task) Insert(task *database.Task, ctx context.Context) (sql.Result, error) {
	return tr.DBConnection.NewInsert().Model(task).Exec(ctx)
}

func (tr Task) Update(task *database.Task, ctx context.Context) (sql.Result, error) {
	return tr.DBConnection.NewUpdate().Model(task).WherePK().Exec(ctx)
}

func (tr Task) Delete(task *database.Task, ctx context.Context) (sql.Result, error) {
	return tr.DBConnection.NewDelete().Model(task).WherePK().Exec(ctx)
}
