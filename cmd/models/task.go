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

type RequestTask struct {
	Text      string `json:"text,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

var NotFoundByIdError = errors.New("not found by ID")

func AddTask(requestTask RequestTask, db *bun.DB) (*database.Task, error) {
	text := strings.TrimSpace(requestTask.Text)

	if utf8.RuneCount([]byte(text)) > 1000 { // ограничение 1000 символов
		return nil, errors.New("text of task too long")
	}

	if utf8.RuneCount([]byte(text)) == 0 {
		return nil, errors.New("text of task is empty")
	}

	task := database.Task{
		Text:      text,
		CreatedAt: time.Now().Unix(),
	}

	if _, err := task.Insert(db, context.Background()); err != nil {
		return nil, err
	}
	return &task, nil
}

func CompleteTask(id int64, db *bun.DB) error {
	ctx := context.Background()

	task, err := database.FindTaskById(id, db, ctx)
	if err == sql.ErrNoRows {
		return NotFoundByIdError
	}
	if err != nil {
		return err
	}

	if task.IsCompleted == true {
		return errors.New("task already completed")
	}

	task.CompletedAt = time.Now().Unix()
	task.IsCompleted = true

	if _, err := task.Update(db, context.Background()); err != nil {
		return err
	}
	return nil
}

func DeleteTask(id int64, db *bun.DB) error {
	ctx := context.Background()

	task, err := database.FindTaskById(id, db, ctx)
	if err == sql.ErrNoRows {
		return NotFoundByIdError
	}
	if err != nil {
		return err
	}

	if _, err := task.Delete(db, context.Background()); err != nil {
		return err
	}
	return nil
}

func GetTask(id int64, db *bun.DB) (*database.Task, error) {
	ctx := context.Background()

	task, err := database.FindTaskById(id, db, ctx)
	if err == sql.ErrNoRows {
		return nil, NotFoundByIdError
	}
	return &task, nil
}
