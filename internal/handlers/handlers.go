package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Sirpyerre/go-search-jobs/pkg/client"
	"net/http"
	"net/url"
	"strings"
)

// IndexRoute index router handler
func IndexRoute(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

// parseParams build query for solr
func parseParams(params map[string][]string) string {
	var query []string
	allowedFields := []string{
		"title",
		"description",
		"salary",
		"location",
		"category",
	}

	if len(params) > 0 {
		format := "%s:%s"
		for _, field := range allowedFields {
			if len(params[field]) > 0 && params[field][0] != "" {
				query = append(query, fmt.Sprintf(format, field, params[field][0]))
			}
		}
	}

	queryStr := strings.Join(query, " OR ")
	return url.QueryEscape(queryStr)
}
