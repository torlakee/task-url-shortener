package main

import (
	"chaos-url-shortener/deployment"
	"chaos-url-shortener/server"
)

func main() {
	deployment.Reloadable()

	app := deployment.StartUp()

	server.Start(app)

	deployment.CleanUp()
}
