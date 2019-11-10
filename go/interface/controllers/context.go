package controllers

// Context controllers interface
type Context interface {
	// Param(string) string
	// Bind(interface{}) error
	// JSON(int, interface{})
	String(int, string) error
}
