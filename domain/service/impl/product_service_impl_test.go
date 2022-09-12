package impl

import (
	entity "go-web-sample/domain/entity"
	repository "go-web-sample/infrastructure/mysql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	service := NewProductService(repository.NewProductRepository())
	new_product := entity.Build("商品-A", 100)
	err := service.Add(new_product)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, true)
}

func TestSearchBynKeyword(t *testing.T) {
	service := NewProductService(repository.NewProductRepository())
	products, err := service.SearchByKeyword("鉛筆")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, len(products), 2)
}
