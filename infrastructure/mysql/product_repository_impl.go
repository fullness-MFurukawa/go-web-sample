package mysql

import (
	"database/sql"
	entity "go-web-sample/domain/entity"
	"go-web-sample/domain/repository"
	"strings"
	"time"
)

const INSERT = "INSERT INTO products (name,price) VALUES (?,?)"
const SELECT_BY_NAME_LIKE = "SELECT id,name,price,register FROM products WHERE name LIKE ?"

// 商品をアクセスするRepositoryインターフェース実装構造体
type productRepository struct {
}

// コンストラクタ
func NewProductRepository() repository.ProductRepository {
	repository := productRepository{}
	return &repository
}

// 指定されたキーワードで商品を部分一致検索する
func (repo *productRepository) SelectByNameLike(tx *sql.Tx, keyword string) ([]*entity.Product, error) {
	stmt, err := tx.Prepare(SELECT_BY_NAME_LIKE)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	var sb strings.Builder
	sb.WriteString("%")
	sb.WriteString(keyword)
	sb.WriteString("%")
	rows, err := stmt.Query(sb.String())
	if err != nil {
		return nil, err
	}
	var products = []*entity.Product{}
	for rows.Next() {
		var id int
		var name string
		var price int
		var register time.Time
		rows.Scan(&id, &name, &price, &register)
		product := entity.NewProduct(id, name, price, register)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// 新しい商品を追加する
func (repo *productRepository) Insert(tx *sql.Tx, product *entity.Product) error {
	stmt, err := tx.Prepare(INSERT)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(product.Name(), product.Price())
	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if num == 0 || err != nil {
		return err
	}
	return nil
}
