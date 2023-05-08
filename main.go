package main

import (
	"myapp/data"
	"myapp/handlers"

	"github.com/hilsonxhero/napoleon"
)

type application struct {
	App      *napoleon.Napoleon
	Handlers *handlers.Handlers
	Models   data.Models
}

func main() {
	init := initApplication()
	init.App.ListenAndServe()
}
