package impl

import (
	"go-web-sample/domain/entity"
	"go-web-sample/domain/repository"
	"go-web-sample/domain/service"
	connector "go-web-sample/infrastructure/mysql"
)

// 商品に対するサービス機能インターフェース実装構造体
type productService struct {
	repository repository.ProductRepository
}

// コンストラクタ
func NewProductService(repository repository.ProductRepository) service.ProductService {
	service := productService{}
	service.repository = repository
	return &service
}

// ProductServiceインターフェースのメソッド実装
// 商品をキーワード検索する
func (sv *productService) SearchByKeyword(keyword string) ([]*entity.Product, error) {
	// データベース接続
	connector := connector.MySqlConnector()
	db, err := connector.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	// トランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	// 商品キーワード検索
	products, err := sv.repository.SelectByNameLike(tx, keyword)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// 新商品を登録する
func (sv *productService) Add(product *entity.Product) error {
	// データベース接続
	connector := connector.MySqlConnector()
	db, err := connector.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	// トランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	err = sv.repository.Insert(tx, product)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
