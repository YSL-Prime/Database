package handler

import (
	"database/entity"
	"database/mysql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewHttpHandler(db *gorm.DB) http.Handler {
	h := &HttpHandler{DB: db}
	mux := mux.NewRouter()

	mux.HandleFunc("/", h.getAllData).Methods(http.MethodGet)
	mux.HandleFunc("/unreviewed", h.getDataUnreviewed).Methods(http.MethodGet)

	return mux
}

type HttpHandler struct {
	DB *gorm.DB
}

// 모든 데이터를 가져옴
func (h *HttpHandler) getAllData(w http.ResponseWriter, r *http.Request) {
	entities, err := h.fetchDataFromDB(nil)
	if err != nil {
		http.Error(w, fmt.Sprint("Server Error: ", err), http.StatusInternalServerError)
		return
	}

	h.sendJSONResponse(&entities, w)
}

// 리뷰하지 않은(OK 값이 null인) 데이터를 가져옴
func (h *HttpHandler) getDataUnreviewed(w http.ResponseWriter, r *http.Request) {
	entities, err := h.fetchDataFromDB(mysql.S2p("ok IS NULL"))
	if err != nil {
		http.Error(w, fmt.Sprint("Server Error: ", err), http.StatusInternalServerError)
		return
	}

	h.sendJSONResponse(&entities, w)
}

// request body에 json으로 하나의 entity.Data가 전달 되면, 해당 레코드의 ok 필드 값을 업데이트 함.

func (h *HttpHandler) fetchDataFromDB(query *string) ([]entity.Data, error) {
	var entities []entity.Data
	if query != nil {
		h.DB = h.DB.Where(*query)
	}

	if result := h.DB.Find(&entities); result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

func (h *HttpHandler) sendJSONResponse(entities *[]entity.Data, w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(entities); err != nil {
		http.Error(w, fmt.Sprint("Server Errer: ", err), http.StatusInternalServerError)
		return
	}
}
