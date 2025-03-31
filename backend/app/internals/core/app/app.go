package app

import (
	"encoding/json"
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
)

// App represents the core application structure, managing routes, middleware, and handling HTTP requests.
type App struct {
	routes     []*Route
	middleware []lib.Middleware
	Hub        *socket.Hub
}

// NewApp creates a new instance of the App with default configurations.
func NewApp() *App {
	return &App{}
}

// Use adds middleware functions to the application middleware stack.
func (a *App) Use(middleware ...lib.Middleware) {
	a.middleware = append(a.middleware, middleware...)
}

// AddRoute adds a route to the framework.
func (a *App) AddRoute(router Route) {
	router.Init()
	a.routes = append(a.routes, &router)
}

// GET is a shorthand method to add a GET route to the framework.
func (a *App) GET(path string, handler lib.Handler) {
	router := Route{
		Path:    path,
		Method:  "GET",
		Handler: handler,
	}
	a.AddRoute(router)
}

// POST is a shorthand method to add a POST route to the framework.
func (a *App) POST(path string, handler lib.Handler) {
	router := Route{
		Path:    path,
		Method:  "POST",
		Handler: handler,
	}
	a.AddRoute(router)
}

// ServeHTTP implements the http.Handler interface to handle incoming HTTP requests.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method
	isRouteFounded := false

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	if method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	for _, route := range a.routes {
		if route.Pattern.MatchString(path) {
			isRouteFounded = true
			if route.Method != method {
				continue
			}
			route.RebuildURLWithParams(r.URL)

			finalHandler := route.Handler
			for _, mw := range a.middleware {
				finalHandler = mw(finalHandler)
			}
			finalHandler(w, r)
			return
		}
	}

	if isRouteFounded {
		a.handleError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}
	a.handleError(w, http.StatusNotFound, "Not Found!")
}

// handleError writes an error response with the specified status code and message.
func (a *App) handleError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Code":    status,
		"Message": message,
		"Data":    nil,
	})
}
