package api

import (
	"net/http"

	"github.com/Dreck2003/api/backend/helpers"
	"github.com/Dreck2003/api/backend/models"
	"github.com/Dreck2003/api/backend/types"
	"github.com/go-chi/chi/v5"
)

var EmailRoutes = []types.RouterType{
	{Pattern: "/{term}", Router: getEmails, TypeMethod: "GET"},
}

func getEmails(w http.ResponseWriter, r *http.Request) {
	stringToSearch := chi.URLParam(r, "term")
	w.Header().Set("Content-Type", "application/json")
	customResponse := helpers.CreateResponse()

	if len(stringToSearch) <= 0 {
		customResponse.AddStatus(400).AddError(types.HttpErrors.NotProvided("term")).Send(&w)
		return
	}

	data, err := models.Email.Search(stringToSearch)

	if err != nil {
		customResponse.AddStatus(500).AddError("An error ocurred").Send(&w)
		return
	}
	customResponse.AddStatus(200).AddContent(data).Send(&w)
}
