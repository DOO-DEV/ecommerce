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

func Register(c echo.Context, client pb.AuthServiceClient) {
	body := RegisterRequestBody{}

	if err := c.Bind(&body); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err)
		return
	}

	res, err := client.Register(c.Request().Context(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		echo.NewHTTPError(int(res.Status), res.Error)
	}

	// I think if pass pointer will be a little performance optimization ;)
	c.JSON(int(res.Status), &res)
}
