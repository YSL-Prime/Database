package main

import (
	"database/handler"
	"database/mysql"
	"fmt"
	"net/http"
)

func main() {
	db, err := mysql.Get()
	if err != nil {
		fmt.Println("DB 연결 실패")
		return
	}
	handler := handler.NewHttpHandler(db)
	http.ListenAndServe(":9999", handler)
}
