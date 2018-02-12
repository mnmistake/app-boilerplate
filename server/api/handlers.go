package api

import (
	"github.com/raunofreiberg/kyrene/server"
	"github.com/raunofreiberg/kyrene/server/model"
	"time"
)

var (
	id          int
	content     string
	isCompleted bool
	createdAt   string
)

var TodoList []model.Todo

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func QueryTodos() (interface{}, error) {
	rows, err := server.DB.Query("SELECT id, content, is_completed, created_at FROM todos ORDER BY created_at")
	todos := TodoList

	checkError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&id,
			&content,
			&isCompleted,
			&createdAt,
		)

		checkError(err)

		todos = append(todos, model.Todo{
			ID:          id,
			Content:     content,
			IsCompleted: isCompleted,
			CreatedAt:   createdAt,
		})
	}

	err = rows.Err()
	checkError(err)

	return todos, nil
}

func QueryTodo(queryID int) (interface{}, error) {
	rows, err := server.DB.Query("SELECT id, content, is_completed, created_at FROM todos WHERE id=$1", queryID)
	checkError(err)

	for rows.Next() {
		err := rows.Scan(&id, &content, &isCompleted, &createdAt)
		checkError(err)

		return model.Todo{
			ID:          id,
			Content:     content,
			IsCompleted: isCompleted,
			CreatedAt:   createdAt,
		}, nil
	}

	err = rows.Err()
	checkError(err)

	panic("No todo found")
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
	_, err := server.DB.Exec(
		"UPDATE todos SET is_completed = $1 WHERE id = $2",
		IsCompleted,
		id,
	)

	if err != nil {
		return nil, err
	}

	return model.Todo{
		ID:          id,
		IsCompleted: IsCompleted,
	}, nil
}

func DeleteTodo(id int) (interface{}, error) {
	_, err := server.DB.Exec("DELETE FROM todos WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	return model.Todo{
		ID: id,
	}, nil
}

func DeleteTodos() (interface{}, error) {
	_, err := server.DB.Exec("DELETE FROM todos *")

	if err != nil {
		return nil, err
	}

	return TodoList, nil
}
