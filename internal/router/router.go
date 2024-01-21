package router

import (
	"log"
	"net/http"

	"theobourgeois.com/app/templates/layout"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

type Vars = map[string]string
type DynamicComponent = func(vars Vars) templ.Component
type Route struct {
	Name      string
	Component templ.Component
}

type DynamicRoute struct {
	Name             string
	ComponentHandler func(vars Vars) templ.Component
}

type ApiRouteHandler = func(w http.ResponseWriter, r *http.Request) templ.Component

type ApiRoute struct {
	Name    string
	Handler ApiRouteHandler
	Method  string
}

var routes []Route
var dynamicRoutes []DynamicRoute
var apiRoutes []ApiRoute

func CreateRoute(name string, component templ.Component) {
	routes = append(routes, Route{name, component})
}

func CreateDynamicRoute(name string, componentHandler func(vars Vars) templ.Component) {
	dynamicRoutes = append(dynamicRoutes, DynamicRoute{name, componentHandler})
}

func CreateApiRoute(name string, method string, handler ApiRouteHandler) {
	apiRoutes = append(apiRoutes, ApiRoute{name, handler, method})
}

func SetupRoutes() {
	r := mux.NewRouter()

	for _, route := range routes {
		r.Handle(route.Name, templ.Handler(layout.Layout(route.Component)))
	}

	for _, dynamicRoute := range dynamicRoutes {
		r.HandleFunc(dynamicRoute.Name, func(w http.ResponseWriter, r *http.Request) {
			params := mux.Vars(r)
			templ.Handler(layout.Layout(dynamicRoute.ComponentHandler(params))).ServeHTTP(w, r)
		})
	}

	for _, apiRoute := range apiRoutes {
		r.HandleFunc(apiRoute.Name, makeApiHandler(apiRoute)).Methods(apiRoute.Method)
	}

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", r)
}

func makeApiHandler(apiRoute ApiRoute) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != apiRoute.Method {
			log.Println("Invalid request method, expected", apiRoute.Method, "got", r.Method, "for", apiRoute.Name, "route")
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		component := apiRoute.Handler(w, r)
		templ.Handler(component).ServeHTTP(w, r)
	}
}
