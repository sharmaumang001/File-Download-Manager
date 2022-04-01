package controller

import (
	"dm/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetFilePath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range services.Direc {
		if item.FileName == params["filename"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
