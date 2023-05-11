package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/hilsonxhero/napoleon"
)

type application struct {
	App        *napoleon.Napoleon
	Handlers   *handlers.Handlers
	Models     data.Models
	MiddleWare *middleware.Middleware
}

func main() {
	init := initApplication()
	init.App.ListenAndServe()
}
