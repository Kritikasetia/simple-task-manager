package app

import (
	"context"
	"database/sql"
)

type Task struct {
	ID          int
	Title       string
	Description string
}

func CreateTask(ctx context.Context, db *sql.DB, task Task) (int, error) {
	result, err := db.ExecContext(ctx, "INSERT INTO tasks (title, description) VALUES (?, ?)", task.Title, task.Description)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func ReadTask(ctx context.Context, db *sql.DB, taskID int) (Task, error) {
	var task Task
	err := db.QueryRowContext(ctx, "SELECT id, title, description FROM tasks WHERE id = ?", taskID).
		Scan(&task.ID, &task.Title, &task.Description)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}
