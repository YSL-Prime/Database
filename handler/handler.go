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

type HttpHandler struct {
	DB *gorm.DB
}

func NewHttpHandler(db *gorm.DB) http.Handler {
	h := &HttpHandler{DB: db}
	mux := mux.NewRouter()

	mux.HandleFunc("/", h.getAllData).Methods(http.MethodGet)
	mux.HandleFunc("/unreviewed", h.getDataUnreviewed).Methods(http.MethodGet)
	mux.HandleFunc("/review/records", h.updateRecords).Methods(http.MethodPost)

	return mux
}

// DB에서 모든 데이터를 가져옴
func (h *HttpHandler) getAllData(w http.ResponseWriter, r *http.Request) {
	entities, err := h.fetchDataFromDB(nil)
	if err != nil {
		http.Error(w, fmt.Sprint("Server Error: ", err), http.StatusInternalServerError)
		return
	}

	h.sendJSONResponse(&entities, w)
}

// DB에서 리뷰하지 않은(OK 값이 null인) 데이터를 가져옴
func (h *HttpHandler) getDataUnreviewed(w http.ResponseWriter, r *http.Request) {
	entities, err := h.fetchDataFromDB(mysql.S2p("ok IS NULL"))
	if err != nil {
		http.Error(w, fmt.Sprint("Server Error: ", err), http.StatusInternalServerError)
		return
	}

	h.sendJSONResponse(&entities, w)
}

// request body에 json으로 하나의 entity.Data가 전달 되면, 해당 레코드의 ok 필드 값을 업데이트 함.
func (h *HttpHandler) updateRecords(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "HTTP POST로 해야함 !!", http.StatusMethodNotAllowed)
		return
	}

	e := new([]entity.Data)
	if err := json.NewDecoder(r.Body).Decode(e); err != nil {
		http.Error(w, fmt.Sprint("Server error: ", err), http.StatusInternalServerError)
		return
	}

	db, err := mysql.Get()
	if err != nil {
		http.Error(w, fmt.Sprint("Server error: ", err), http.StatusInternalServerError)
		return
	}

	if result := db.Save(&e); result.Error != nil {
		http.Error(w, fmt.Sprint("Server error: ", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// 실제 DB에 접근하여 값을 가져오는 메소드
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

// JSON 형태로 response body를 구성하는 메소드
func (h *HttpHandler) sendJSONResponse(entities *[]entity.Data, w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(entities); err != nil {
		http.Error(w, fmt.Sprint("Server Errer: ", err), http.StatusInternalServerError)
		return
	}
}
