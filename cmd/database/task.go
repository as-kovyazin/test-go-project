package database

type Task struct {
	ID          int64 `bun:",pk,autoincrement"`
	CreatedAt   int64
	Text        string
	IsCompleted bool
	CompletedAt int64
}
