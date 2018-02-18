package api

import (
	"errors"
	"time"

	"github.com/raunofreiberg/kyrene/server"
	"github.com/raunofreiberg/kyrene/server/model"
)

var (
	id          int
	content     string
	isCompleted bool
	createdAt   string
)

var TodoList []model.Todo

func QueryTodos() (interface{}, error) {
	rows, err := server.DB.Query("SELECT id, content, is_completed, created_at FROM todos ORDER BY created_at")
	todos := TodoList

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id, &content, &isCompleted, &createdAt); err != nil {
			return nil, err
		}

		todos = append(todos, model.Todo{
			ID:          id,
			Content:     content,
			IsCompleted: isCompleted,
			CreatedAt:   createdAt,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func QueryTodo(queryID int) (interface{}, error) {
	rows, err := server.DB.Query("SELECT id, content, is_completed, created_at FROM todos WHERE id=$1", queryID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&id, &content, &isCompleted, &createdAt); err != nil {
			return nil, err
		}

		return model.Todo{
			ID:          id,
			Content:     content,
			IsCompleted: isCompleted,
			CreatedAt:   createdAt,
		}, nil
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return nil, errors.New("No todo found")
}

func InsertTodo(content string) (interface{}, error) {
	currTime := time.Now().Local()
	err := server.DB.QueryRow(
		"INSERT INTO todos (content, is_completed, created_at) VALUES ($1, $2, $3) RETURNING id",
		content,
		false,
		currTime,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return model.Todo{
		ID:          id,
		Content:     content,
		IsCompleted: false, // todos are marked as uncompleted by default
		CreatedAt:   currTime.String(),
	}, nil
}

func UpdateTodo(id int, IsCompleted bool) (interface{}, error) {
	rows, err := server.DB.Query(
		"UPDATE todos SET is_completed = $1 WHERE id = $2 RETURNING content",
		IsCompleted,
		id,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&content); err != nil {
			return nil, err
		}

		return model.Todo{
			ID:          id,
			IsCompleted: IsCompleted,
			Content:	 content,
		}, nil
	}

	return nil, errors.New("Todo not updated")
}

func DeleteTodo(id int) (interface{}, error) {
	rows, err := server.DB.Query("DELETE FROM todos WHERE id = $1 RETURNING id", id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}

		return model.Todo{}, nil
	}

	return nil, errors.New("Todo not deleted")
}

func DeleteTodos() (interface{}, error) {
	_, err := server.DB.Exec("DELETE FROM todos *")

	if err != nil {
		return nil, err
	}

	return TodoList, nil
}
