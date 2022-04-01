package main

import (
	"dm/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.Router(r)
	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
