package main

import (
	"github.com/JuanigTorres/xmen-finder/database"
	"log"
	"net/http"

	"github.com/JuanigTorres/xmen-finder/controller"
)

// TODO:
//  - Subir a Heroku
func main() {
	database.NewClient()
	defer database.Disconnect()
	http.HandleFunc("/mutant", controller.MutantHandler)
	http.HandleFunc("/stats", controller.StatsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
