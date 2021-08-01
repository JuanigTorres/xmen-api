package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JuanigTorres/xmen-api/controller"
	"github.com/JuanigTorres/xmen-api/database"
)

// TODO:
//  - Agregar README.md
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9290"
	}

	database.NewClient()
	defer database.Disconnect()
	http.HandleFunc("/", http.RedirectHandler("/stats", http.StatusMovedPermanently).ServeHTTP)
	http.HandleFunc("/mutant", controller.MutantHandler)
	http.HandleFunc("/stats", controller.StatsHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
