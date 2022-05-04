package main

import (
	"fmt"
	"github.com/Sirpyerre/go-search-jobs/internal/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const portNumber = ":3000"

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.IndexRoute).Methods("GET")
	router.HandleFunc("/facet", handlers.FacetRoute).Methods("GET")
	fmt.Println(fmt.Sprintf("Starting application in port %q", portNumber))
	log.Fatal(http.ListenAndServe(portNumber, router))
}
