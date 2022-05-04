package helpers

import (
	"fmt"
	"log"
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

// ParseParams build query for solr
func ParseParams(params map[string][]string) string {
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

func ParseFacetField(params map[string][]string) string {
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
