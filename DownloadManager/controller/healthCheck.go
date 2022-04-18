package controller

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Msg string `json:"message"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := Message{
		Msg: "Server is running at port no 8000",
	}
	json.NewEncoder(w).Encode(&err.Msg)
}
