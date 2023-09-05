package main

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"api-gateway/pkg/order"
	"api-gateway/pkg/product"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	e := echo.New()

	authSvc := *auth.RegisterRoutes(e, &c)

	product.RegisterRoutes(e, &c, &authSvc)
	order.RegisterRoutes(e, &c, &authSvc)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", c.Port)))
}
