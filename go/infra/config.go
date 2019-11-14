package infra

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Specification application enviroment variables
type Specification struct {
	Port int `default:"1323"`
}

// AppEnv application environment variables
var AppEnv *Specification

// EnvSet setting env
func EnvSet() {
	var s Specification
	err := envconfig.Process("app", &s)

	if err != nil {
		log.Fatal(err.Error())
	}

	AppEnv = &s
}
