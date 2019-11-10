package infra

import (
	"github.com/Lim79Plus/go-postgres/go/interface/controllers"
	"github.com/labstack/echo"
)

// Router *echo.Echo
var Router *echo.Echo

func init() {
	router := echo.New()

	router.GET("/", func(c echo.Context) error {
		return controllers.SayHello(c)
	})

	Router = router
}
