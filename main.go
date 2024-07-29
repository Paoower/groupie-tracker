package main

import (
	"groupie-tracker/handlers"
	"groupie-tracker/utils"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Ltime)
	log.SetPrefix("groupie-tracker:")

	go getData()

	scripts := http.FileServer(http.Dir("./templates/scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", scripts))

	http.HandleFunc("/searchbar", handlers.SearchBarHandler)
	http.HandleFunc("/filters", handlers.FiltersHandler)
	http.HandleFunc("/artists", handlers.ArtistsHandler)
	http.HandleFunc("/artist", handlers.ArtistHandler)
	http.HandleFunc("/", handlers.HomeHandler)

	log.Println(utils.GREEN, "Server started at http://localhost:8080", utils.NONE)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getData() {
	err := utils.ReadJson("json/artists.json", &utils.ArtistsJson)
	if err != nil {
		log.Fatal(err)
	}

	err = utils.ReadJson("json/locations.json", &utils.LocationsJson)
	if err != nil {
		log.Fatal(utils.RED, err)
	}

	err = utils.ReadJson("json/relation.json", &utils.RelationJson)
	if err != nil {
		log.Fatal(utils.RED, err)
	}

	log.Println(utils.GREEN, "Data loaded", utils.NONE)
}
