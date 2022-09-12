package presentation

import (
	entity "go-web-sample/domain/entity"
)

// 商品情報
type ProductJson struct {
	Id       int    `json:"商品番号"`
	Name     string `json:"商品名"`
	Price    int    `json:"単価"`
	Register string `json:"登録日"`
}

// 複数の商品情報JSON形式に変換
func ProductsConverter(products []*entity.Product) []*ProductJson {
	jsonData := make([]*ProductJson, 0, len(products))
	for _, value := range products {
		row := &ProductJson{Id: value.Id(), Name: value.Name(), Price: value.Price(), Register: value.Register().String()}
		jsonData = append(jsonData, row)
	}
	return jsonData
}
