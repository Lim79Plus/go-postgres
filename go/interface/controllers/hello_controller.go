package controllers

import (
	"net/http"
)

// HelloController return say hello
type HelloController interface {
	Hello(c Context)
}

// SayHello return hello for you
func SayHello(c Context) error {
	return c.String(http.StatusOK, "Hello Workd!\n")
}
