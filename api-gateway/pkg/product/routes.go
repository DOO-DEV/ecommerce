package product

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"api-gateway/pkg/product/routes"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := e.Group("/product", a.AuthRequired())
	routes.POST("", svc.CreateProduct)
	routes.GET("/:id", svc.FindOne)
}

func (svc *ServiceClient) FindOne(c echo.Context) error {
	return routes.FindOne(c, svc.Client)
}

func (svc *ServiceClient) CreateProduct(c echo.Context) error {
	return routes.CreateProduct(c, svc.Client)
}
