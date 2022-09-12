package repository

import (
	"database/sql"
	entity "go-web-sample/domain/entity"
)

// 商品をアクセスするRepositoryインターフェース
type ProductRepository interface {
	// 新しい商品を追加する
	Insert(tx *sql.Tx, product *entity.Product) error
	// 指定されたキーワードで商品を部分一致検索する
	SelectByNameLike(tx *sql.Tx, keyword string) ([]*entity.Product, error)
}
