package mysql

import (
	"database/entity"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	assert := assert.New(t)

	db, err := Get()
	assert.NoError(err)

	err = db.Migrator().DropTable(&entity.Data{})
	assert.NoError(err)

	err = db.AutoMigrate(&entity.Data{})
	assert.NoError(err)

	for i := 0; i < 5; i++ {
		e := &entity.Data{
			Address:  "유성구 봉명동",
			Name:     S2p(fmt.Sprintf("%d번 째 가로등", i+1)),
			Ds:       fmt.Sprintf("유성구 봉명동에 있는 %d번 째 가로등이다.", i+1),
			Category: S2p("편의시설"),
			ImageUrl: fmt.Sprintf("대충 %d번 째 image url", i+1),
		}
		result := db.Create(e)
		assert.NoError(result.Error)
	}

	var cnt int64
	db.Model(&entity.Data{}).Count(&cnt)
	assert.Equal(int64(5), cnt)
}
