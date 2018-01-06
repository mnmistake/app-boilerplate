package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

//"time"

var db *sql.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func dbConfig() map[string]string {
	conf := make(map[string]string)
	conf[dbhost] = dbhost
	conf[dbport] = dbport
	conf[dbuser] = dbuser
	conf[dbpass] = dbpass
	conf[dbname] = dbname

	for key, _ := range conf {
		val, exists := os.LookupEnv(key)
		if !exists {
			panic("oh no")
		}
		conf[key] = val
	}
	return conf
}

func initDb() {
	config := dbConfig()
	var err error

	psqlInfo := fmt.Sprintf(
		`
		host=%s 
		port=%s 
		user=%s
		password=%s
		dbname=%s
		sslmode=disable
		`,
		config[dbhost],
		config[dbport],
		config[dbuser],
		config[dbpass],
		config[dbname],
	)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB =>", config[dbname])
}

func main() {
	initDb()
	defer db.Close()
	http.HandleFunc("/api/index/", indexHandler)
	http.HandleFunc("/api/data/", dataHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	//
}
