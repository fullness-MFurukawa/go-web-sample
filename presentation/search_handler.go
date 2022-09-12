package presentation

import (
	"fmt"
	service "go-web-sample/domain/service/impl"
	repository "go-web-sample/infrastructure/mysql"
	"net/http"

	"github.com/labstack/echo/v4"
)

//
// リクエストパラメータ
//
type SearchParam struct {
	Keyword string `json:"keyword" xml:"keyword" form:"keyword" query:"keyword"`
}

// Get:/search
func Search(c echo.Context) error {
	searchparam := new(SearchParam)
	if err := c.Bind(searchparam); err != nil {
		return err
	}
	service := service.NewProductService(repository.NewProductRepository())
	products, err := service.SearchByKeyword(searchparam.Keyword)

	if err != nil {
		return err
	}
	if len(products) == 0 {
		return c.JSON(http.StatusCreated,
			fmt.Sprintf("キーワード:%sが含まれる商品は見つかりませんでした。", searchparam.Keyword))
	}
	productsJson := ProductsConverter(products)
	return c.JSON(http.StatusCreated, productsJson)
}
