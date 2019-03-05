package api

import (
	"net/http"
	. "travel-n-expenses/api/clients"
	. "travel-n-expenses/api/login"
	. "travel-n-expenses/api/reports"

	"github.com/gorilla/mux"
)

const staticDir = "/api/assets/"

//Here there are all the routes from all modules
func AddAllRoutes() *mux.Router {
	router := mux.NewRouter()
	// Server CSS, JS & Images Statically.
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	// Reports Routes
	AddReportsHandler(router)
	AddClientsHandler(router)
	AddLoginHandler(router)
	return router
}
