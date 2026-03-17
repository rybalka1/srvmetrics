package main

import (
	"net/http"

	"github.com/rybalka1/srvmetrics/internal/handlers"
	"github.com/rybalka1/srvmetrics/internal/storage"
)

func main() {
	st := storage.NewMemStorage()

	h := handlers.NewHandler(st)
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, h.UpdateMetric)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
