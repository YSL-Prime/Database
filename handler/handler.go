package handler

import (
	"database/entity"
	"database/mysql"
	"encoding/json"
	"fmt"
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
	entities, err := fetchDataFromDB(nil)
	if err != nil {
		http.Error(w, fmt.Sprint("Server Error: ", err), http.StatusInternalServerError)
		return
	}

	sendJSONResponse(&entities, w)
}

// 리뷰하지 않은(OK 값이 null인) 데이터를 가져옴
func getDataUnreviewed(w http.ResponseWriter, r *http.Request) {
	entities, err := fetchDataFromDB(mysql.S2p("ok IS NULL"))
	if err != nil {
		http.Error(w, fmt.Sprint("Server Error: ", err), http.StatusInternalServerError)
		return
	}

	sendJSONResponse(&entities, w)
}

func fetchDataFromDB(query *string) ([]entity.Data, error) {
	db, err := mysql.Get()
	if err != nil {
		return nil, err
	}

	var entities []entity.Data
	if query != nil {
		db = db.Where(*query)
	}

	if result := db.Find(&entities); result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

func sendJSONResponse(entities *[]entity.Data, w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(entities); err != nil {
		http.Error(w, fmt.Sprint("Server Errer: ", err), http.StatusInternalServerError)
		return
	}
}
