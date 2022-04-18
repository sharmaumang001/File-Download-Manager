package main

import (
	"DWM/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	Router := mux.NewRouter()
	routes.Router(Router)
	fmt.Println("server started at port no 8000")
	log.Fatal(http.ListenAndServe(":8000", Router))

}
