package main

import (
	"github.com/hector-leite/meli-notification/application"
	"github.com/hector-leite/meli-notification/server"
)

func main() {
	server.Run(application.BuildApp())
}
