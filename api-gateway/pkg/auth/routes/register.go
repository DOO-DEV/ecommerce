package routes

import (
	"api-gateway/pkg/auth/pb"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c echo.Context, client pb.AuthServiceClient) error {
	body := RegisterRequestBody{}

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := client.Register(c.Request().Context(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		return echo.NewHTTPError(int(res.Status), res.Error)
	}

	// I think if pass pointer will be a little performance optimization ;)
	return c.JSON(int(res.Status), &res)
}
