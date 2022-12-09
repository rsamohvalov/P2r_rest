/*
 * Сервис интерфейса P2r
 *
 * Сервис для реализации процедур интерфейса P2r. 
 *
 * API version: 1.0.
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"ReleaseConnection",
		strings.ToUpper("Delete"),
		"/connections/{rm_id}",
		ReleaseConnection,
	},

	Route{
		"SetupConnection",
		strings.ToUpper("Post"),
		"/connections",
		SetupConnection,
	},

	Route{
		"DeleteWarning",
		strings.ToUpper("Delete"),
		"/connections/{rm_id}/warnings/{warning_id}",
		DeleteWarning,
	},

	Route{
		"SendWarning",
		strings.ToUpper("Post"),
		"/connections/{rm_id}/warnings",
		SendWarning,
	},
}