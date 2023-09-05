package routes

import (
	"api-gateway/pkg/auth/pb"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c echo.Context, client pb.AuthServiceClient) error {
	body := LoginRequestBody{}

	if err := c.Bind(&body); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err)
		return err
	}

	res, err := client.Login(c.Request().Context(), &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		echo.NewHTTPError(int(res.Status), res.Error)
	}

	c.JSON(int(res.Status), &body)

	return nil
}
