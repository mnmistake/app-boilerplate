package main

import (
	"fmt"
	"log"
)

func QueryTodos() error {
	rows, err := db.Query(`SELECT id, content FROM todos`)

	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		//todos := Todo{}
		err := rows.Scan(
			&id,
			&content,
		)
		if err != nil {
			return err
		}
		fmt.Println(id, content)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func QueryTodo(queryID int) interface{} {
	data, err := db.Query("SELECT id, content, is_completed FROM todos WHERE id=$1", queryID)
	if err != nil {
		log.Fatal(err)
	}

	for data.Next() {
		err := data.Scan(&id, &content, &isCompleted)
		if err != nil {
			log.Fatal(err)
		}

		return Todo{
			ID:          id,
			Content:     content,
			IsCompleted: isCompleted,
		}
	}

	err = data.Err()
	if err != nil {
		log.Fatal(err)
	}
	return Todo{}
}

/*func DataHandler(w http.ResponseWriter, r *http.Request) {
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
*/
