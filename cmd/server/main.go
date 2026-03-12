package main

import (
	"github.com/rybalka1/srvmetrics/internal/handlers"
	"github.com/rybalka1/srvmetrics/internal/storage"
	"net/http"
)

func main() {
	storage := storage.NewMemStorage()
	handlers.SetStorage(storage)

	mux := http.NewServeMux()
	mux.HandleFunc(`/api`, handlers.ApiPage)
	mux.HandleFunc(`/`, handlers.MainPage)
	mux.HandleFunc(`/update/`, handlers.UpdateMetric)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
