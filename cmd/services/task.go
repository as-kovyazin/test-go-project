package services

import (
	"context"
	"database/sql"
	"errors"
	"iSpringTest/database"
	"iSpringTest/repositories"
	"strings"
	"time"
	"unicode/utf8"
)

type Task struct {
	Repository *repositories.Task
}

type RequestTask struct {
	Text      string `json:"text,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

var NotFoundByIdError = errors.New("not found by ID")
var DatabaseError = errors.New("database error")

func (t Task) AddTask(requestTask RequestTask) (*database.Task, error) {
	text := strings.TrimSpace(requestTask.Text)
	cntOfSymbols := utf8.RuneCount([]byte(text))

	if cntOfSymbols > 1000 { // ограничение 1000 символов
		return nil, errors.New("text of task too long")
	}

	if cntOfSymbols == 0 {
		return nil, errors.New("text of task is empty")
	}

	task := database.Task{
		Text:      text,
		CreatedAt: time.Now().Unix(),
	}

	if _, err := t.Repository.Insert(&task, context.Background()); err != nil {
		return nil, DatabaseError
	}
	return &task, nil
}

func (t Task) CompleteTask(id int64) error {
	ctx := context.Background()

	task, err := t.Repository.FindTaskById(id, ctx)
	if err == sql.ErrNoRows {
		return NotFoundByIdError
	}
	if err != nil {
		return DatabaseError
	}

	if task.IsCompleted == true {
		return errors.New("task already completed")
	}

	task.CompletedAt = time.Now().Unix()
	task.IsCompleted = true

	if _, err := t.Repository.Update(&task, context.Background()); err != nil {
		return DatabaseError
	}
	return nil
}

func (t Task) DeleteTask(id int64) error {
	ctx := context.Background()

	task, err := t.Repository.FindTaskById(id, ctx)
	if err == sql.ErrNoRows {
		return NotFoundByIdError
	}
	if err != nil {
		return DatabaseError
	}

	if _, err := t.Repository.Delete(&task, context.Background()); err != nil {
		return DatabaseError
	}
	return nil
}

func (t Task) GetTask(id int64) (*database.Task, error) {
	ctx := context.Background()

	task, err := t.Repository.FindTaskById(id, ctx)
	if err == sql.ErrNoRows {
		return nil, NotFoundByIdError
	}
	if err != nil {
		return nil, DatabaseError
	}
	return &task, nil
}
