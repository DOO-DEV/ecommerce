package routes

import (
	"api-gateway/pkg/product/pb"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Sku   string `json:"sku"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"` // TODO - must be a float64 in here and proto request
}

func CreateProduct(c echo.Context, client pb.ProductServiceClient) error {
	b := CreateProductRequestBody{}

	if err := c.Bind(&b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	req := &pb.CreateProductRequest{
		Name:  b.Name,
		Sku:   b.Sku,
		Stock: b.Stock,
		Price: b.Price,
	}

	res, err := client.CreateProduct(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(int(res.Status), res.Error)
	}

	return c.JSON(int(res.Status), &res)
}
