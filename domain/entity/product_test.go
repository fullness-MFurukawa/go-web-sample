package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetId(t *testing.T) {
	product := Product{}
	err1 := product.SetId(0)
	assert.EqualError(t, err1, "無効なIdです。")
	err2 := product.SetId(-1)
	assert.EqualError(t, err2, "無効なIdです。")
	err3 := product.SetId(100)
	assert.Nil(t, err3)
}

func TestSetName(t *testing.T) {
	product := Product{}
	err1 := product.SetName("")
	assert.EqualError(t, err1, "商品名は必須です。")
	err2 := product.SetName("あいうえおかきくけこさしすせそた")
	assert.EqualError(t, err2, "商品名は15文字以内です。")
	err3 := product.SetName("商品-ABC")
	assert.Nil(t, err3)
}

func TestSetPrice(t *testing.T) {
	product := Product{}
	err1 := product.SetPrice(0)
	assert.EqualError(t, err1, "無効な単価です。")
	err2 := product.SetPrice(-100)
	assert.EqualError(t, err2, "無効な単価です。")
	err3 := product.SetPrice(40)
	assert.EqualError(t, err3, "単価は50以上、5000以下です。")
	err4 := product.SetPrice(5500)
	assert.EqualError(t, err4, "単価は50以上、5000以下です。")
	err5 := product.SetPrice(3000)
	assert.Nil(t, err5)
}
