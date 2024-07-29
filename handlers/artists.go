package handlers

import (
	"errors"
	"groupie-tracker/datatypes"
	"groupie-tracker/utils"
	"html/template"
	"log"
	"net/http"
)

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	data := datatypes.ArtistsPage{
		Artists:   utils.ArtistsJson,
		Locations: utils.GetAllLocations(utils.LocationsJson),
	}

	t, err := template.ParseFiles("templates/artists.html", "templates/artist_block.html")
	if err != nil {
		log.Println(utils.RED, "Error parsing artist template", utils.NONE)
		ErrorHandler(w, errors.New("500 | Internal server error: Could not parse template"))
		return
	}
	t.Execute(w, data)
}
