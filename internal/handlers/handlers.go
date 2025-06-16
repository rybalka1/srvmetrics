package handlers

import "net/http"

type MyHandler struct{}

func MainPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Привет!"))
}

func ApiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Это страница /api."))
}
