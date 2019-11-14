package infra

import (
	"github.com/Lim79Plus/go-postgres/go/interface/controllers"
	"github.com/labstack/echo"
)

// Router *echo.Echo
var Router *echo.Echo

func init() {
	// set env
	EnvSet()

	router := echo.New()

	router.GET("/", handler(controllers.SayHello))

	Router = router
}
