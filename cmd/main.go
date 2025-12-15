package main

import "rianRestapp/routes"

type AppConfig struct {
	Port string
}

func main() {
	data := &AppConfig{
		Port: ":6060",
	}
	routes.IntialRoute(data.Port)

}
