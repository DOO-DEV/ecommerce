package routes

import (
	"api-gateway/pkg/order/pb"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateOrderRequestBody struct {
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(c echo.Context, client pb.OrderServiceClient) error {
	b := CreateOrderRequestBody{}

	if err := c.Bind(&b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	userID := c.Get("userID").(int64)

	res, err := client.CreateOrder(c.Request().Context(), &pb.CreateOrderRequest{
		ProductId: b.ProductID,
		Quantity:  b.Quantity,
		UserId:    userID,
	})
	if err != nil {
		return echo.NewHTTPError(int(res.Status), res.Error)
	}

	return c.JSON(http.StatusCreated, &res)
}
