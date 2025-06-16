package main

import (
	"github.com/rybalka1/srvmetrics/internal/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/api`, handlers.ApiPage)
	mux.HandleFunc(`/`, handlers.MainPage)
	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
