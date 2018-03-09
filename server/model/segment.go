package model

type Segment struct {
	ID        int    `json:"id,omitempty"`
	SheetID   int    `json:"sheetId,omitempty"`
	Label     string `json:"label,omitempty"`
	Content   string `json:"content,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}
