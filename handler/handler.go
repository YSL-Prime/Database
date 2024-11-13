package handler

import (
	"database/entity"
	"database/mysql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHttpHandler() http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/", getAllData).Methods(http.MethodGet)
	mux.HandleFunc("/unreviewed", getDataUnreviewed).Methods(http.MethodGet)

	return mux
}

// 모든 데이터를 가져옴
func getAllData(w http.ResponseWriter, r *http.Request) {
	db, err := mysql.Get()
	if err != nil {
		http.Error(w, "서버 오류: DB 연결이 실패했습니다.", http.StatusInternalServerError)
		return
	}

	var entities []entity.Data
	if result := db.Find(&entities); result.Error != nil {
		http.Error(w, "서버 오류: DB에서 데이터를 찾는 데 실패했습니다.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(entities); err != nil {
		http.Error(w, "서버 오류: 데이터를 JSON으로 변환하는 데 실패했습니다.", http.StatusInternalServerError)
		return
	}
}

// 리뷰하지 않은(OK 값이 null인) 데이터를 가져옴
func getDataUnreviewed(w http.ResponseWriter, r *http.Request) {
	db, err := mysql.Get()
	if err != nil {
		http.Error(w, "서버 오류: DB 연결이 실패했습니다.", http.StatusInternalServerError)
		return
	}

	var entities []entity.Data
	if result := db.Where("ok IS NULL").Find(&entities); result.Error != nil {
		http.Error(w, "서버 오류: DB에서 데이터를 찾는 데 실패했습니다.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(entities); err != nil {
		http.Error(w, "서버 오류: 데이터를 JSON으로 변환하는 데 실패했습니다.", http.StatusInternalServerError)
		return
	}
}
