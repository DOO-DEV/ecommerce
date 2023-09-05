package order

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"api-gateway/pkg/order/routes"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := e.Group("/order", a.AuthRequired())
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(c echo.Context) error {
	return routes.CreateOrder(c, svc.Client)
}
