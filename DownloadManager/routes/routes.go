package routes

import (
	"dm/controller"

	"github.com/gorilla/mux"
)

var Router = func(R *mux.Router) {

	R.HandleFunc("/download", controller.DownloadFileSeq).Methods("POST")
	R.HandleFunc("/filepath/{id}", controller.GetFilePath).Methods("GET")
	R.HandleFunc("/healthcheck", controller.HealthCheck).Methods("GET")
}
