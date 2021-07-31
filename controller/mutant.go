package controller

import (
	"encoding/json"
	"net/http"

	"github.com/JuanigTorres/xmen-finder/model"
	"github.com/JuanigTorres/xmen-finder/service"
)

func MutantHandler(response http.ResponseWriter, request *http.Request) {
	status := http.StatusForbidden
	switch request.Method {
	case http.MethodPost:
		var data model.XmenRequest
		if err := json.NewDecoder(request.Body).Decode(&data); err == nil {
			if isMutant, ex := service.IsMutant(data.DNA); ex == nil && isMutant {
				status = http.StatusOK
			}
		}
	default:
		status = http.StatusMethodNotAllowed
	}
	response.WriteHeader(status)
}
