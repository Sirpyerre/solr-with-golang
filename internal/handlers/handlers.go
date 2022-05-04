package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Sirpyerre/go-search-jobs/internal/helpers"
	"github.com/Sirpyerre/go-search-jobs/pkg/client"
	"net/http"
)

// IndexRoute index router handler
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	params := helpers.ParseParams(vars)
	if params == "" {
		http.Error(w, "missing params", http.StatusBadRequest)
		return
	}

	query := "query?q="
	docs := client.GetQuery("gettingstarted", query+params)
	result, err := json.Marshal(docs.Response.Docs)
	if err != nil {
		fmt.Println("error marshal:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(result)
	if err != nil {
		http.Error(w, "something is wrong", http.StatusBadRequest)
		return
	}
}

// FacetRoute facet router handler
func FacetRoute(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	query := "query?q=" + helpers.ParseParams(vars)
	query += "&facet=true&" + helpers.ParseFacetField(vars)
	docs := client.GetQuery("gettingstarted", query)
	result, err := json.Marshal(docs.FacetCounts["facet_fields"])
	if err != nil {
		fmt.Println("error marshal:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(result)
	if err != nil {
		http.Error(w, "something is wrong", http.StatusBadRequest)
		return
	}
}
