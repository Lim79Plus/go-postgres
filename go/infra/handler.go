package infra

import (
	"github.com/Lim79Plus/go-postgres/go/interface/controllers"
	"github.com/labstack/echo"
)

func handler(f func(c controllers.Context) error) func(c echo.Context) error {
	return func(c echo.Context) error {
		return f(c)
	}
}
