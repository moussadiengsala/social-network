package main

import (
	"fmt"
	"net/http"

	core "learn.zone01dakar.sn/forum-rest-api/app/internals/core/app"
	"learn.zone01dakar.sn/forum-rest-api/app/routes"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/middlewares"
	ws "learn.zone01dakar.sn/forum-rest-api/app/service/socket"
)

func main() {

	// Creating a new instance of our mini-framework
	app := core.NewApp()

	// Middleware usage
	app.Use(service.LoggerMiddleware)
	app.Use(service.AuthMiddleware)
	app.Use(service.DBMiddleware)

	// call different endpoint handlers.
	routes.Handle(app)

	// websocket
	ws.Ws(app)

	fmt.Println("server runnin at http://localhost:8000")
	// Serving our app
	http.ListenAndServe(":8000", app)
}
