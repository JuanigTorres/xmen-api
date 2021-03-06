package controller

import (
	"encoding/json"
	"github.com/JuanigTorres/xmen-api/database"
	"github.com/JuanigTorres/xmen-api/model/documents"
	"net/http"

	"github.com/JuanigTorres/xmen-api/model"
	"github.com/JuanigTorres/xmen-api/service"
)

func MutantHandler(response http.ResponseWriter, request *http.Request) {
	status := http.StatusForbidden
	switch request.Method {
	case http.MethodPost:
		var data model.XmenRequest
		if err := json.NewDecoder(request.Body).Decode(&data); err == nil {
			if isMutant, ex := service.IsMutant(data.DNA); ex == nil && isMutant {
				database.SaveDNA(documents.NewDNADocument(data.DNA, isMutant))
				status = http.StatusOK
			}
		}
	default:
		status = http.StatusMethodNotAllowed
	}
	response.WriteHeader(status)
}
