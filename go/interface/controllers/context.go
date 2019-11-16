package controllers

// Context controllers interface
type Context interface {
	JSON(int, interface{}) error
}
