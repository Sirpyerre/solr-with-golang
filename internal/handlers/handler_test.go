package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var theTest = []struct {
	name               string
	url                string
	method             string
	params             string
	expectedStatusCode int
}{
	{"home", "/xsdsdsd", "GET", "", 404},
	{"search", "/?", "GET", "title=golang&category=fulltime", http.StatusOK},
	{"facets", "/facet?", "GET", "field=location", http.StatusOK},
}

func TestHandlers(t *testing.T) {
	ts := httptest.NewTLSServer(Routers())
	defer ts.Close()

	for _, test := range theTest {
		log.Println("testUrl:", ts.URL+test.url+test.params)
		resp, err := ts.Client().Get(ts.URL + test.url + test.params)
		if err != nil {
			t.Log(err)
			log.Fatal(err)
		}

		if resp.StatusCode != test.expectedStatusCode {
			t.Errorf("for %s, expected %d but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
		}
	}
}

func Routers() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", IndexRoute).Methods("GET")
	router.HandleFunc("/facet", FacetRoute).Methods("GET")

	return router
}
