package models

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/Dreck2003/api/backend/constants"
	"github.com/Dreck2003/api/backend/helpers"
	"github.com/Dreck2003/api/backend/types"
)

type EmailsModel struct {
	types.ModelSearcher
}

type QueryT struct {
	Term string `json:"term"`
}

type SearchStruct struct {
	SearchType string   `json:"search_type"`
	Query      QueryT   `json:"query"`
	SortFields []string `json:"sort_fields"`
	From       uint64   `json:"from"`
	MaxResults uint64   `json:"max_results"`
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

var URL = constants.Env[constants.ZINC_URL] + "api/emails/_search"

var Email = EmailsModel{}

func (e *EmailsModel) Search(term string) (*ResponseZincApi, error) {
	parseJson, err := json.Marshal(SearchStruct{
		SearchType: "match",
		Query: QueryT{
			Term: term,
		},
		From:       0,
		MaxResults: 20,
	})

	if err != nil {
		return nil, err
	}

	req := helpers.CreateRequest("POST").SetUrl(URL).Build(bytes.NewBuffer(parseJson))
	req.HttpRequest.SetBasicAuth(constants.Env[constants.ZINC_USERNAME], constants.Env[constants.ZINC_PASSWORD])
	response, err := req.Send()

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseStruct ResponseZincApi
	parseError := json.Unmarshal(body, &responseStruct)

	if parseError != nil {
		return nil, parseError
	}

	return &responseStruct, nil
}
