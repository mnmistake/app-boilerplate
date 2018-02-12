package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"	
)

func LoginFunc(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	fmt.Println(decoder)
}

func RegisterFunc(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	userData := struct {
		Username	string	`json:"username"`
		Password	string	`json:"password"`
	}{}
	err := decoder.Decode(&userData)

	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	fmt.Println(userData)
}