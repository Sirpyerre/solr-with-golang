package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Sirpyerre/go-search-jobs/pkg/client"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var allowedFields = []string{
	"title",
	"description",
	"salary",
	"location",
	"category",
}

// IndexRoute index router handler
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	params := parseParams(vars)
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

	//if params == "" {
	//	http.Error(w, "missing params", http.StatusBadRequest)
	//	return
	//}

	query := "query?q=" + parseParams(vars)
	query += "&facet=true&" + parseFacetField(vars)
	docs := client.GetQuery("gettingstarted", query)
	result, err := json.Marshal(docs.FacetCounts["facet_fields"])
	if err != nil {
		fmt.Println("error marshal:", err)
	}

	log.Println("facetCounts:", docs.FacetCounts["facet_fields"])

	facetFields := docs.FacetCounts["facet_fields"].(map[string]interface{})
	for i, field := range facetFields {
		log.Println("i: ", i, "value:", field)
		for _, val := range field.([]interface{}) {
			log.Println("val:", val)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(result)
	if err != nil {
		http.Error(w, "something is wrong", http.StatusBadRequest)
		return
	}
}

// parseParams build query for solr
func parseParams(params map[string][]string) string {
	var query []string
	if len(params) > 0 {
		format := "%s:%s"
		for _, field := range allowedFields {
			if len(params[field]) > 0 && params[field][0] != "" {
				query = append(query, fmt.Sprintf(format, field, params[field][0]))
			}
		}
	}

	if len(query) == 0 {
		query = append(query, "*:*")
	}
	log.Println("query:", query)

	queryStr := strings.Join(query, " OR ")
	return url.QueryEscape(queryStr)
}

func parseFacetField(params map[string][]string) string {
	log.Println("facet params", params)
	query := ""
	if len(params) > 0 {
		format := "%s=%s"
		for _, field := range allowedFields {
			if len(params["field"]) > 0 && params["field"][0] == field {
				query = fmt.Sprintf(format, "facet.field", params["field"][0])
			}
		}
	}

	log.Println("query facet:", query)
	return query
}
