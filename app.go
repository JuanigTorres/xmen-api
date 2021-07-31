package main

import (
	"github.com/JuanigTorres/xmen-finder/database"
	"log"
	"net/http"

	"github.com/JuanigTorres/xmen-finder/controller"
)

// TODO:
//  - Probar paralelismo
//  - Subir a Heroku
//  - Falta servicio de /stats
func main()  {
	database.NewClient()
	defer database.Disconnect()
	http.HandleFunc("/mutant", controller.MutantHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}