package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JuanigTorres/xmen-finder/controller"
	"github.com/JuanigTorres/xmen-finder/database"
)

// TODO:
//  - Comentar
//  - Agregar README.md
//  - Rename go.mod
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9290"
	}

	database.NewClient()
	defer database.Disconnect()
	http.HandleFunc("/mutant", controller.MutantHandler)
	http.HandleFunc("/stats", controller.StatsHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
