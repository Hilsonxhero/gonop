package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"

	"github.com/hilsonxhero/napoleon"
)

func initApplication() *application {

	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	napoleonApp := &napoleon.Napoleon{}
	err = napoleonApp.New(path)
	if err != nil {
		log.Fatal(err)
	}

	napoleonApp.AppName = "myApp"

	myMiddleWare := &middleware.Middleware{
		App: napoleonApp,
	}

	myHandlers := &handlers.Handlers{
		App: napoleonApp,
	}

	app := &application{
		App:        napoleonApp,
		Handlers:   myHandlers,
		MiddleWare: myMiddleWare,
	}

	app.Models = data.New(app.App.DB.Pool)

	myHandlers.Models = app.Models
	myMiddleWare.Models = app.Models
	app.App.Routes = app.routes()

	return app
}
