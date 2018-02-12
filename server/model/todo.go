package model

type Todo struct {
	ID          int    `json:"id,omitempty"`
	Content     string `json:"content,omitempty"`
	IsCompleted bool   `json:"isCompleted,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
}
