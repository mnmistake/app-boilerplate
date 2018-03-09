package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

type User struct {
	ID       int `gorm:"PRIMARY_KEY"`
	Username string
	Password []uint8
}

func InitDb() {
	var err error

	psqlInfo := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s sslmode=disable",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
	)

	DB, err = gorm.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to DB =>", os.Getenv("DB_NAME"))

	DB.LogMode(true)

	DB.AutoMigrate(&User{})

	//DB.Exec("CREATE TABLE IF NOT EXISTS todos (id SERIAL PRIMARY KEY, content TEXT, is_completed BOOL, created_at TEXT)")
	//DB.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username TEXT, password TEXT)")
}
