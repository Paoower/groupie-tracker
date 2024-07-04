package handlers

import (
	"errors"
	"groupie-tracker/datatypes"
	"groupie-tracker/utils"
	"html/template"
	"net/http"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	requestID := r.URL.Query().Get("id")

	var a datatypes.ArtistResponse
	artistID, _ := strconv.Atoi(requestID)
	a.Artist = utils.GetArtist(artistID, utils.ArtistsJson)
	a.Relations = utils.GetRelations(artistID, utils.RelationJson)

	t, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		ErrorHandler(w, errors.New("500 | Internal server error: Could not parse template"))
	}
	t.Execute(w, a)
}
