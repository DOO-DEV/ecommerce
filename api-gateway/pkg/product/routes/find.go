package routes

import (
	"api-gateway/pkg/product/pb"
	"github.com/labstack/echo/v4"
)

func FindOne(c echo.Context, client pb.ProductServiceClient) error {
	id := c.Get("userID").(int64)

	res, err := client.FindOne(c.Request().Context(), &pb.FindOneRequest{Id: id})
	if err != nil {
		return echo.NewHTTPError(int(res.Status), res.Error)
	}

	return c.JSON(int(res.Status), &res)
}
