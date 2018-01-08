package main

var (
	id          int
	content     string
	isCompleted bool
)

func QueryTodos() interface{} {
	rows, err := db.Query("SELECT id, content, is_completed FROM todos ORDER BY id")
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

func InsertTodo(content string) interface{} {
	err := db.QueryRow("INSERT INTO todos (content, is_completed) VALUES ($1, $2) RETURNING id", content, false).Scan(&id)
	checkError(err)

	return Todo{
		ID:          id,
		Content:     content,
		IsCompleted: false,
	}
}

func UpdateTodo(id int, content string, IsCompleted bool) interface{} {
	err := db.QueryRow("UPDATE todos SET content = $1, is_completed = $2 WHERE id = $3 RETURNING id", content, isCompleted, id).Scan(&id)
	checkError(err)

	return Todo{
		ID:          id,
		Content:     content,
		IsCompleted: IsCompleted,
	}
}

func DeleteTodo(id int) interface{} {
	_, err := db.Exec("DELETE FROM todos WHERE id = $1", id)
	checkError(err)

	return Todo{
		ID: id,
	}
}
