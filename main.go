package main

import (
	"database/handler"
	"net/http"
)

func main() {
	handler := handler.NewHttpHandler()
	http.ListenAndServe(":9999", handler)
}
