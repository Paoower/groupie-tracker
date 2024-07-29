package main

import (
	"groupie-tracker/handlers"
	"groupie-tracker/utils"
	"log"
	"net/http"
	"os"
)

func init() {
	log.SetFlags(log.Ltime)
	log.SetPrefix("groupie-tracker:")
	getData()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/searchbar":
		handlers.SearchBarHandler(w, r)
	case "/filters":
		handlers.FiltersHandler(w, r)
	case "/artists":
		handlers.ArtistsHandler(w, r)
	case "/artist":
		handlers.ArtistHandler(w, r)
	case "/":
		handlers.HomeHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is ready to handle requests at :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func getData() {
	if err := utils.ReadJson("json/artists.json", &utils.ArtistsJson); err != nil {
		log.Printf("Error loading artists.json: %v", err)
	}
	if err := utils.ReadJson("json/locations.json", &utils.LocationsJson); err != nil {
		log.Printf("Error loading locations.json: %v", err)
	}
	if err := utils.ReadJson("json/relation.json", &utils.RelationJson); err != nil {
		log.Printf("Error loading relation.json: %v", err)
	}
	log.Println("Data loaded successfully")
}
