package main

import "time"

type Todo struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Completed bool      `json:"completed,omitempty"`
	Due       time.Time `json:"due,omitempty"`
}

type Todos []Todo
