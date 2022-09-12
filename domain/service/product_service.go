package service

import (
	entity "go-web-sample/domain/entity"
)

// 商品に対するサービス機能インターフェース
type ProductService interface {
	// 新商品を追加する
	Add(product *entity.Product) error
	// 商品をキーワード検索する
	SearchByKeyword(keyword string) ([]*entity.Product, error)
}
