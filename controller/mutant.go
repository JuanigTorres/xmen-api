package controller

import (
	"encoding/json"
	"github.com/JuanigTorres/xmen-finder/database"
	"github.com/JuanigTorres/xmen-finder/model/documents"
	"net/http"

	"github.com/JuanigTorres/xmen-finder/model"
	"github.com/JuanigTorres/xmen-finder/service"
)

func MutantHandler(response http.ResponseWriter, request *http.Request) {
	status  := http.StatusForbidden
	switch request.Method {
	case http.MethodPost:
		var data model.XmenRequest
		if err := json.NewDecoder(request.Body).Decode(&data); err == nil {
			isMutant, ex := service.IsMutant(data.DNA)
			if ex == nil {
				if isMutant {
					status = http.StatusOK
				}
				database.SaveDNA(documents.NewDNADocument(data.DNA, isMutant))
			}
		}
	default:
		status = http.StatusMethodNotAllowed
	}
	response.WriteHeader(status)
}
