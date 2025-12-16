package main

import (
	"os"
	"rianRestapp/routes"
)

type AppConfig struct {
	Port          string
	ApiVersioning string
}

func main() {
	data := &AppConfig{
		Port:          ":6060",
		ApiVersioning: os.Getenv("APP_VERSIONING"),
	}
	routes.IntialRoute(data.Port)

}
