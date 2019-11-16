package controllers

import (
	"net/http"
)

// HelloController return say hello
type HelloController interface {
	Hello(c Context)
}

// Resp is respose body
type Resp struct {
	Msg string `json:"msg"`
}

// SayHello return hello for you
func SayHello(c Context) error {
	m := Resp{
		Msg: "Hello Workd!",
	}
	return c.JSON(http.StatusOK, m)
}
