package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ResponseSolr struct {
	ResponseHeader ResponseHeader `json:"responseHeader"`
	Response       Response       `json:"response"`
}

type ResponseHeader struct {
	Status int `json:"status"`
	QTime  int `json:"QTime"`
	Params struct {
		Q string `json:"q"`
	} `json:"params"`
}

type Response struct {
	NumFound      int    `json:"numFound"`
	Start         int    `json:"start"`
	NumFoundExact bool   `json:"numFoundExact"`
	Docs          []Docs `json:"docs"`
}

type Docs struct {
	Title       []string `json:"title"`
	Description []string `json:"description"`
	Salary      []string `json:"salary"`
	Category    []string `json:"category"`
	ID          string   `json:"id"`
	Version     int64    `json:"_version_"`
}

const urlBase = "http://localhost:8983/solr"

func GetQuery(core, query string) []Docs {
	format := "%s/%s/query?q=%s"
	url := fmt.Sprintf(format, urlBase, core, query)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject ResponseSolr
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		fmt.Println("unmarshal error:", err)
	}

	return responseObject.Response.Docs
}
