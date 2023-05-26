package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/uptrace/bun"
	"iSpringTest/database"
	"strings"
	"time"
	"unicode/utf8"
)

type RequestTodo struct {
	Text string `json:"text"`
}

func AddTodo(requestTodo RequestTodo, db *bun.DB) (*database.Todo, error) {
	text := strings.TrimSpace(requestTodo.Text)

	if utf8.RuneCount([]byte(text)) > 1000 { // ограничение 1000 символов
		return nil, errors.New("text of task too long")
	}

	if utf8.RuneCount([]byte(text)) == 0 {
		return nil, errors.New("text of task is empty")
	}

	dbTodo := database.Todo{
		Text:      text,
		CreatedAt: time.Now().Unix(),
	}

	if _, err := dbTodo.Insert(db, context.Background()); err != nil {
		return nil, err
	}
	return &dbTodo, nil
}

func CompleteTodo(id int64, db *bun.DB) error {
	ctx := context.Background()

	todo, err := database.FindTodoById(id, db, ctx)
	if err == sql.ErrNoRows {
		return errors.New("not found by ID")
	}
	if err != nil {
		return err
	}

	if todo.IsCompleted == true {
		return errors.New("task already completed")
	}

	todo.CompletedAt = time.Now().Unix()
	todo.IsCompleted = true

	if _, err := todo.Update(db, context.Background()); err != nil {
		return err
	}
	return nil
}
