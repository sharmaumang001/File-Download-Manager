package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func FilePath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	json.NewEncoder(w).Encode("/Users/vikash/Desktop/DownloadManager/app/" + params["id"])

}
