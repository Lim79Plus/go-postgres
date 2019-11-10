package infra

import (
	"fmt"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// Specification application enviroment variables
type Specification struct {
	Port int
}

// AppEnv application environment variables
var AppEnv *Specification

// EnvSet setting env
func EnvSet() {
	var s Specification
	err := envconfig.Process("app", &s)
	fmt.Println("EnvSet app s", s)
	if err != nil {
		log.Fatal(err.Error())
	}

	AppEnv = &s

	// test
	fmt.Println("1. $APP_PORT:", os.Getenv("APP_PORT"))
	fmt.Println("1. $USER:", os.Getenv("USER"))
}
