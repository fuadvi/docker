package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/about", aboutHandler)

	fmt.Println("server started at localhost:8080")
	http.ListenAndServe(":8080", r)
}

type Response struct {
	Messegae string      `json:"message"`
	Data     interface{} `json:"data"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	Response := Response{
		Messegae: "Ok",
		Data: map[string]string{
			"description": "This is about page",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response)
}
