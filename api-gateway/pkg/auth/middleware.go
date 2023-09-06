package auth

import (
	"api-gateway/pkg/auth/pb"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type MiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) MiddlewareConfig {
	return MiddlewareConfig{svc: svc}
}

func (m *MiddlewareConfig) AuthRequired() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			jwtToken := strings.Split(authorization, "Bearer ")
			if len(jwtToken) < 2 {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			res, err := m.svc.Client.Validate(c.Request().Context(), &pb.ValidateRequest{Token: jwtToken[1]})
			if err != nil {
				return echo.NewHTTPError(int(res.Status), res.Error)
			}

			c.Set("userID", res.UserId)
			return next(c)
		}
	}
}
