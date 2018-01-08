package main

var (
	id          int
	content     string
	isCompleted bool
)

func QueryTodos() interface{} {
	rows, err := db.Query("SELECT id, content, is_completed FROM todos")
	todos := TodoList

	checkError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&id,
			&content,
			&isCompleted,
		)
		todos = append(todos, Todo{
			ID:          id,
			Content:     content,
			IsCompleted: isCompleted,
		})
		checkError(err)
	}

	err = rows.Err()
	checkError(err)

	return todos
}

func QueryTodo(queryID int) interface{} {
	rows, err := db.Query("SELECT id, content, is_completed FROM todos WHERE id=$1", queryID)
	checkError(err)

	for rows.Next() {
		err := rows.Scan(&id, &content, &isCompleted)
		checkError(err)

		return Todo{
			ID:          id,
			Content:     content,
			IsCompleted: isCompleted,
		}
	}

	err = rows.Err()
	checkError(err)

	return Todo{}
}
