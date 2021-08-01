package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JuanigTorres/xmen-api/service"
)

func StatsHandler(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		res, err := json.Marshal(service.GetStats())
		if err == nil {
			_, err = response.Write(res)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Println("WARN ::", err)
		}
	default:
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
}
