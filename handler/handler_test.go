package handler

import (
	"bytes"
	"database/entity"
	"database/mysql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 실제 DB에 의존함
// 가져온 데이터를 검증하지 않고 데이터의 개수를 하드코딩하여 검증함
func TestGetAllData(t *testing.T) {
	assert := assert.New(t)

	// HTTP 요청 생성
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	// 핸들러 실행
	db, err := mysql.Get()
	assert.Nil(err)

	handler := NewHttpHandler(db)
	handler.ServeHTTP(res, req)

	// 응답 코드 확인
	assert.Equal(http.StatusOK, res.Code)

	// 응답 본문(JSON) 확인
	var responseDatas []entity.Data
	err = json.NewDecoder(res.Body).Decode(&responseDatas)
	assert.Nil(err)

	// Mock 데이터와 응답이 동일한지 확인
	assert.Equal(5, len(responseDatas))
}

func TestGetDataUnreviewed(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest(http.MethodGet, "/unreviewed", nil)
	res := httptest.NewRecorder()

	db, err := mysql.Get()
	assert.Nil(err)

	handler := NewHttpHandler(db)
	handler.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	var responseDatas []entity.Data
	err = json.NewDecoder(res.Body).Decode(&responseDatas)
	assert.Nil(err)
	assert.Equal(4, len(responseDatas))
}

func TestUpdateRecord(t *testing.T) {
	assert := assert.New(t)

	// 확인하지 않은 값 가져오기
	req := httptest.NewRequest(http.MethodGet, "/unreviewed", nil)
	res := httptest.NewRecorder()

	db, err := mysql.Get()
	assert.Nil(err)

	h := NewHttpHandler(db)
	h.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	var datas []entity.Data
	err = json.NewDecoder(res.Body).Decode(&datas)
	assert.Nil(err)

	// 값 업데이트
	d := datas[0:2]
	for i := range d {
		v := true
		d[i].Ok = &v
	}

	jsonData, err := json.Marshal(d)
	assert.Nil(err)
	req = httptest.NewRequest(http.MethodPost, "/review/records", bytes.NewReader(jsonData))
	res = httptest.NewRecorder()

	h.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
}
