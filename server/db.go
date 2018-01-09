package server

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

//"time"

var DB *sql.DB

/*const (
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

*/

func InitDb() {	
	var err error

	shit := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s sslmode=disable",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
	)
	
	fmt.Println(os.Getenv("POSTGRES_PASSWORD"))
	fmt.Println(shit)

	DB, err = sql.Open("postgres", shit)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Connected to DB =>", os.Getenv("DB_NAME"))

	DB.Exec("CREATE TABLE IF NOT EXISTS todos (id SERIAL PRIMARY KEY, content TEXT, is_completed BOOL)")
}
