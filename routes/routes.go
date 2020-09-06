package routes

import (
	filmcontroller "film-service/controllers/film"
	pingcontroller "film-service/controllers/ping"
	loggingmiddleware "film-service/middleware/logger"
	"net/http"

	"github.com/gorilla/mux"
)

// DefineRoutes setup the routes for the service
func DefineRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingmiddleware.Middleware)
	r.HandleFunc("/ping", pingcontroller.Ping).Methods(http.MethodGet)
	r.HandleFunc("/upgraded-ping", pingcontroller.UpgradedPing).Methods(http.MethodGet)

	r.HandleFunc("/get-films", filmcontroller.GetFilms).Methods(http.MethodGet)
	r.HandleFunc("/create-films", filmcontroller.CreateFilms).Methods(http.MethodPost)
	return r
}
