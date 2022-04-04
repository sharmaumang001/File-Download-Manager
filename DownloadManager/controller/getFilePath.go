package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetFilePath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	json.NewEncoder(w).Encode("/Users/aashutoshkashyap/Desktop/DownloadManager/app/" + params["id"])
}
