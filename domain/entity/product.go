package entity

import (
	"errors"
	"time"
	"unicode/utf8"
)

// 商品ドメインオブジェクト
type Product struct {
	id       int       //	商品番号
	name     string    //	商品名
	price    int       //	単価
	register time.Time //	登録日
}

// ゲッター
func (p *Product) Id() int             { return p.id }
func (p *Product) Name() string        { return p.name }
func (p *Product) Price() int          { return p.price }
func (p *Product) Register() time.Time { return p.register }

// セッター
func (p *Product) SetId(id int) error {
	if id <= 0 {
		return errors.New("無効なIdです。")
	}
	p.id = id
	return nil
}
func (p *Product) SetName(name string) error {
	if name == "" {
		return errors.New("商品名は必須です。")
	}
	if utf8.RuneCountInString(name) > 15 {
		return errors.New("商品名は15文字以内です。")
	}
	p.name = name
	return nil
}
func (p *Product) SetPrice(price int) error {
	if price <= 0 {
		return errors.New("無効な単価です。")
	}
	if price < 50 || price > 5000 {
		return errors.New("単価は50以上、5000以下です。")
	}
	p.price = price
	return nil
}
func (p *Product) SetRegister(register time.Time) {
	p.register = register
}

// コンストラクタ
func NewProduct(id int, name string, price int, register time.Time) *Product {
	product := new(Product)
	product.SetId(id)
	product.SetName(name)
	product.SetPrice(price)
	product.SetRegister(register)
	return product
}
func Build(name string, price int) *Product {
	product := new(Product)
	product.SetName(name)
	product.SetPrice(price)
	return product
}
