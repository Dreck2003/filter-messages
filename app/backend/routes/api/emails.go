package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Dreck2003/api/backend/helpers"
	"github.com/Dreck2003/api/backend/types"
	"github.com/go-chi/chi/v5"
)

var URL = "http://localhost:4080/api/emails/_search"

var EmailRoutes = []types.RouterType{
	{Pattern: "/{term}", Router: getEmails, TypeMethod: "GET"},
}

type ResponseHitsZincApi struct {
	Index  string `json:"_index"`
	Id     string `json:"_id"`
	Source struct {
		Content string `json:"content"`
		EmailId string `json:"emailId"`
		From    string `json:"from"`
		Subject string `json:"subject"`
		To      string `json:"to"`
	} `json:"_source"`
}

type ResponseZincApi struct {
	Hits struct {
		Total struct {
			Value uint `json:"value"`
		} `json:"total"`
		Hits []ResponseHitsZincApi `json:"hits"`
	} `json:"hits"`
}

func insertSearch(term string) string {
	return fmt.Sprintf(`{
		"search_type": "match",
    	"query": {
        "term": "%s"
    	},
    	"sort_fields": ["-@timestamp"],
    	"from": 0,
    	"max_results": 20
		}`,
		term,
	)
}

func getEmails(w http.ResponseWriter, r *http.Request) {
	stringToSearch := chi.URLParam(r, "term")
	w.Header().Set("Content-Type", "application/json")
	customResponse := helpers.CreateResponse()

	if len(stringToSearch) <= 0 {
		customResponse.AddStatus(400).AddError(types.HttpErrors.NotProvided("term")).Send(&w)
		return
	}
	request, err := http.NewRequest("POST", URL, strings.NewReader(insertSearch(stringToSearch)))
	if err != nil {
		customResponse.AddStatus(500).AddError("An error ocurred").Send(&w)
		return
	}
	request.SetBasicAuth("username", "password")
	response, err := http.DefaultClient.Do(request) // Send the http request
	if err != nil {
		customResponse.AddStatus(500).AddError("An error ocurred").Send(&w)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		customResponse.AddStatus(500).AddError("An error ocurred").Send(&w)
		return
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		customResponse.AddStatus(500).AddError("An error ocurred").Send(&w)
		return
	}

	var responseStruct ResponseZincApi
	parseError := json.Unmarshal(body, &responseStruct)

	if parseError != nil {
		customResponse.AddStatus(500).AddError("An error ocurred").Send(&w)
		return
	}
	customResponse.AddStatus(200).AddContent(responseStruct).Send(&w)
}
