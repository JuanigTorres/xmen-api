package main

import (
	"log"
	"net/http"

	"github.com/JuanigTorres/xmen-finder/controller"
	"github.com/JuanigTorres/xmen-finder/database"
)

// TODO:
//  - Comentar
//  - Agregar README.md
//  - Rename go.mod
func main() {
	database.NewClient()
	defer database.Disconnect()
	http.HandleFunc("/mutant", controller.MutantHandler)
	http.HandleFunc("/stats", controller.StatsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
