package model

type Sheet struct {
	ID        int    `json:"id,omitempty"`
	UserID    int    `json:"userId,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}
