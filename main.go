package main

import (
	"log"
	"net/http"
)

func main() {
	router := InitRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
