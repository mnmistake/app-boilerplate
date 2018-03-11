package database

// All database models
var Models = []interface{}{
	&User{},
	&Sheet{},
	&Segment{},
}

type User struct {
	ID       int
	Username string
	Password []uint8
	Sheets   []*Sheet
}

type Sheet struct {
	ID        int
	UserID    int
	Name      string
	CreatedAt string
	Segments  []*Segment
}

type Segment struct {
	ID        int
	SheetID   int
	Label     string
	Content   string
	CreatedAt string
}
