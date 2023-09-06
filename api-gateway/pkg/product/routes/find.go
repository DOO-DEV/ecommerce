package routes

import (
	"api-gateway/pkg/product/pb"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func FindOne(c echo.Context, client pb.ProductServiceClient) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	res, err := client.FindOne(c.Request().Context(), &pb.FindOneRequest{Id: id})
	if err != nil {
		return echo.NewHTTPError(int(res.Status), res.Error)
	}

	return c.JSON(int(res.Status), &res)
}
