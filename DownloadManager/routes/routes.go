package routes

import (
	"DWM/controller"

	"github.com/gorilla/mux"
)

var Router = func(r *mux.Router) {
	r.HandleFunc("/DownloadFile", controller.DownloadFiles).Methods("POST")
	r.HandleFunc("/FilePath/{id}", controller.FilePath).Methods("GET")
	r.HandleFunc("/HealthCheck", controller.HealthCheck).Methods("GET")

}
