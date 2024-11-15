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
	err = http.ListenAndServe("127.0.0.1:9999", handler)
	if err != nil {
		fmt.Println("서버 실행 오류: ", err)
	}
}
