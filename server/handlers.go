package server

var (
	id          int
	content     string
	isCompleted bool
)

type Todo struct {
	ID          int    `json:"id,omitempty"`
	Content     string `json:"content,omitempty"`
	IsCompleted bool   `json:"isCompleted,omitempty"`
}

var TodoList []Todo

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

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

		checkError(err)

		todos = append(todos, Todo{
			ID:          id,
			Content:     content,
			IsCompleted: isCompleted,
		})
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

	panic("No todo found")
}

func InsertTodo(content string) interface{} {
	err := db.QueryRow(
		"INSERT INTO todos (content, is_completed) VALUES ($1, $2) RETURNING id",
		content,
		false,
	).Scan(&id)
	checkError(err)

	return Todo{
		ID:          id,
		Content:     content,
		IsCompleted: false,
	}
}

func UpdateTodo(id int, content string, IsCompleted bool) interface{} {
	_, err := db.Exec(
		"UPDATE todos SET content = $1, is_completed = $2 WHERE id = $3",
		content,
		isCompleted,
		id,
	)
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

func DeleteTodos() interface{} {
	db.Exec("DELETE FROM todos *")

	return TodoList
}
