package auth

import (
	"api-gateway/pkg/auth/routes"
	"api-gateway/pkg/config"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := e.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(c echo.Context) error {
	return routes.Register(c, svc.Client)
}

func (svc *ServiceClient) Login(c echo.Context) error {
	return routes.Login(c, svc.Client)
}
