package handler

import (
	"database/entity"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 실제 데이터의 값을 비교하지는 않음
// 그냥 Test DB에 있는 데이터 수와 일치하는 지만 하드코딩으로 비교함
func TestGetAllData(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	handler := NewHttpHandler()
	handler.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	var responseDatas []entity.Data
	err := json.NewDecoder(res.Body).Decode(&responseDatas)
	assert.Nil(err)
	assert.Equal(5, len(responseDatas))
}

func TestGetDataUnreviewed(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest(http.MethodGet, "/Unreviewed", nil)
	res := httptest.NewRecorder()

	handler := NewHttpHandler()
	handler.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	var responseDatas []entity.Data
	err := json.NewDecoder(res.Body).Decode(&responseDatas)
	assert.Nil(err)
	assert.Equal(4, len(responseDatas))
}
