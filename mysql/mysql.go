package mysql

import (
	"database/entity"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "jaehyeok34:3400@tcp(127.0.0.1:3306)/yuseong?charset=utf8mb4&parseTime=True&loc=Local"

func Get() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db 연결 문제 발생:", err)
		return nil, err
	}

	if err := db.AutoMigrate(&entity.Data{}); err != nil {
		fmt.Println("db 테이블 연결(생성) 문제 발생:", err)
		return nil, err
	}

	return db, nil
}

func S2p(s string) *string {
	return &s
}
