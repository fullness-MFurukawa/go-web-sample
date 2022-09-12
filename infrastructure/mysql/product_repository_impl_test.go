package mysql

import (
	"fmt"
	entity "go-web-sample/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 商品キーワード検索 テストケース1
func TestSelectByNameLike1(t *testing.T) {
	// データベース接続
	connector := MySqlConnector()
	db, _ := connector.Connect()
	defer db.Close()
	// トランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		t.Error(err)
	}
	repository := NewProductRepository()
	products, err := repository.SelectByNameLike(tx, "ABC")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 0, len(products))
}

// 商品キーワード検索 テストケース2
func TestSelectByNameLike2(t *testing.T) {
	// データベース接続
	connector := MySqlConnector()
	db, _ := connector.Connect()
	defer db.Close()
	// トランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		t.Error(err)
	}
	repository := NewProductRepository()
	products, err := repository.SelectByNameLike(tx, "ペン")
	if err != nil {
		t.Error(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}
	assert.Equal(t, 7, len(products))
}

func TestInsert(t *testing.T) {
	// データベース接続
	connector := MySqlConnector()
	db, _ := connector.Connect()
	defer db.Close()
	// トランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		t.Error(err)
	}
	defer tx.Rollback()
	repository := NewProductRepository()
	new_product := entity.Build("商品-A", 100)
	err = repository.Insert(tx, new_product)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, true)
}
