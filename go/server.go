package main

import "github.com/Lim79Plus/go-postgres/go/infra"

func main() {
	infra.Router.Start(":1323")
}
