package handlers

import (
	"encoding/json"
	"errors"
	"groupie-tracker/datatypes"
	"groupie-tracker/utils"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func FiltersHandler(w http.ResponseWriter, r *http.Request) {
	request := datatypes.FiltersRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		ErrorHandler(w, errors.New("500 | Internal server error: Could not decode json"))
		return
	}

	var filtered []datatypes.Artist

	for _, artist := range utils.ArtistsJson {
		match := true

		if artist.CreationDate < request.FromCreationYear {
			match = false
		}

		if artist.CreationDate > request.ToCreationYear {
			match = false
		}

		firstAlbum, _ := strconv.Atoi(strings.Split(artist.FirstAlbum, "-")[2])

		if firstAlbum < request.FromFirstAlbumYear {
			match = false
		}

		if firstAlbum > request.ToFirstAlbumYear {
			match = false
		}

		if request.MembersMin == request.MembersMax && len(artist.Members) != request.MembersMin {
			match = false
		}

		if len(artist.Members) < request.MembersMin || len(artist.Members) > request.MembersMax {
			match = false
		}

		// Check location
		if request.Location != "" {
			locations := utils.GetLocations(artist.Id)
			locationMatch := false
			for _, location := range locations {
				if strings.Contains(strings.ToLower(location), strings.ToLower(request.Location)) {
					locationMatch = true
					break
				}
			}
			if !locationMatch {
				match = false
			}
		}

		if match {
			filtered = append(filtered, artist)
		}
	}

	t, err := template.ParseFiles("templates/artist_block.html")
	if err != nil {
		ErrorHandler(w, errors.New("500 | Internal server error: Could not parse template"))
		return
	}
	t.ExecuteTemplate(w, "artist_block", filtered)
}
