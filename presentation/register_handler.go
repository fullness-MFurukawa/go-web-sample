package presentation

import (
	"fmt"
	entity "go-web-sample/domain/entity"
	service "go-web-sample/domain/service/impl"
	repository "go-web-sample/infrastructure/mysql"
	"net/http"

	"github.com/labstack/echo/v4"
)

//
// リクエストパラメータ
//
type RegisterParam struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Price int    `json:"price" xml:"price" form:"price" query:"price"`
}

// Post:/register
func Register(c echo.Context) error {
	registerparam := new(RegisterParam)
	if err := c.Bind(registerparam); err != nil {
		return err
	}
	service := service.NewProductService(repository.NewProductRepository())
	product := entity.Build(registerparam.Name, registerparam.Price)
	err := service.Add(product)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated,
		fmt.Sprintf("商品:%sを登録しました。", registerparam.Name))
}
