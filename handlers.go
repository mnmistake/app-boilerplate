package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type todo struct {
	ID          int
	Name        string
	IsCompleted bool
}

type todos struct {
	Todos []todo
}

func queryTodos(todos *todos) error {
	rows, err := db.Query(`SELECT * FROM todos`)

	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		todo := todo{}
		err = rows.Scan(
			&todo.ID,
			&todo.Name,
			&todo.IsCompleted,
		)
		if err != nil {
			return err
		}
		todos.Todos = append(todos.Todos, todo)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	todos := todos{}
	err := queryTodos(&todos)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(output))

}
