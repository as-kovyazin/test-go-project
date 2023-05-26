package models

import (
	"context"
	"github.com/uptrace/bun"
	"iSpringTest/database"
)

type AddTodoType struct {
	Text string `json:"text"`
}

func AddTodo(todo AddTodoType, db *bun.DB) error {
	dbTodo := database.Todo{
		Text: todo.Text,
	}

	if _, err := dbTodo.Insert(db, context.Background()); err != nil {
		return err
	}
	return nil
}
