package server

import (
	"time"
)

var (
	id          int
	content     string
	isCompleted bool
	createdAt   string
)

type Todo struct {
	ID          int    `json:"id,omitempty"`
	Content     string `json:"content,omitempty"`
	IsCompleted bool   `json:"isCompleted,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
}

var TodoList []Todo

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func QueryTodos() interface{} {
	rows, err := DB.Query("SELECT id, content, is_completed, created_at FROM todos ORDER BY created_at")
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

		todos = append(todos, Todo{
			ID:          id,
			Content:     content,
			IsCompleted: isCompleted,
			CreatedAt:   createdAt,
		})
	}

	err = rows.Err()
	checkError(err)

	return todos
}

func QueryTodo(queryID int) interface{} {
	rows, err := DB.Query("SELECT id, content, is_completed, created_at FROM todos WHERE id=$1", queryID)
	checkError(err)

	for rows.Next() {
		err := rows.Scan(&id, &content, &isCompleted, &createdAt)
		checkError(err)

		return Todo{
			ID:          id,
			Content:     content,
			IsCompleted: isCompleted,
			CreatedAt:   createdAt,
		}
	}

	err = rows.Err()
	checkError(err)

	panic("No todo found")
}

func InsertTodo(content string) interface{} {
	currTime := time.Now().Local()
	err := DB.QueryRow(
		"INSERT INTO todos (content, is_completed, created_at) VALUES ($1, $2, $3) RETURNING id",
		content,
		false,
		currTime,
	).Scan(&id)
	checkError(err)

	return Todo{
		ID:          id,
		Content:     content,
		IsCompleted: false, // todos are marked as uncompleted by default
		CreatedAt:   currTime.String(),
	}
}

func UpdateTodo(id int, IsCompleted bool) interface{} {
	_, err := DB.Exec(
		"UPDATE todos SET is_completed = $1 WHERE id = $2",
		IsCompleted,
		id,
	)
	checkError(err)

	return Todo{
		ID:          id,
		IsCompleted: IsCompleted,
	}
}

func DeleteTodo(id int) interface{} {
	_, err := DB.Exec("DELETE FROM todos WHERE id = $1", id)
	checkError(err)

	return Todo{
		ID: id,
	}
}

func DeleteTodos() interface{} {
	_, err := DB.Exec("DELETE FROM todos *")
	checkError(err)

	return TodoList
}
