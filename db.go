package main

var currentId int
var todos Todos

func init() {
	CreateTodo(
		Todo{
			Name: "Go shopping",
		},
	)
	CreateTodo(
		Todo{
			Name: "Do shit",
		},
	)
}

func FindTodo(id int) Todo {
	for _, todo := range todos {
		if todo.ID == id {
			return todo // can we remove this ?
		}
	}
	return Todo{}
}

func CreateTodo(t Todo) Todo {
	currentId++
	t.ID = currentId
	todos = append(todos, t)
	return t
}

func DeleteTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return nil
}
