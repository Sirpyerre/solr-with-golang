package main

import (
	"encoding/json"
	"fmt"
	"github.com/Sirpyerre/go-search-jobs/pkg/client"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

const portNumber = ":3000"

func indexRoute(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")

	query := parseParams(vars)
	if query == "" {
		http.Error(w, "missing params", http.StatusBadRequest)
		return
	}

	docs := client.GetQuery("gettingstarted", query)
	result, err := json.Marshal(docs)
	if err != nil {
		fmt.Println("error marshal:", err)
	}

	log.Println(string(result))

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func parseParams(params map[string][]string) string {
	query := ""
	fmt.Println("params", params, "len", len(params))
	if len(params) > 0 {
		searchTerm := params["search"][0]
		if searchTerm != "" {
			format := "title:%s OR description:%s OR salary:%s OR location:%s OR category:%s"
			query = fmt.Sprintf(format, searchTerm, searchTerm, searchTerm, searchTerm, searchTerm)
		}
	}

	return url.QueryEscape(query)
}

func main() {
	fmt.Println("Hola API")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute).Methods("GET")

	fmt.Println(fmt.Sprintf("Starting application in port %q", portNumber))
	log.Fatal(http.ListenAndServe(portNumber, router))
}
