package database

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	"github.com/raunofreiberg/kyrene/server"
)

func createSchema(db *pg.DB) error {
	for _, model := range Models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func Database() *pg.DB {
	port := server.GetEnvWithDefault(os.Getenv("DB_PORT"), "5432")
	addr := fmt.Sprintf("%v:%v", os.Getenv("DB_HOST"), port)

	return pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
		Addr:     addr,
	})
}

func Migrate() {
	db := Database()
	err := createSchema(db)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migrating... => ", os.Getenv("DB_NAME"))
}
