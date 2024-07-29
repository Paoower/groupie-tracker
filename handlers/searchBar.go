package handlers

import (
	"encoding/json"
	"errors"
	"groupie-tracker/datatypes"
	"groupie-tracker/utils"
	"net/http"
	"strconv"
	"strings"
)

func SearchBarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorHandler(w, errors.New("400 | Wrong method"))
		return
	}

	var s datatypes.SearchRequest

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		ErrorHandler(w, errors.New("500 | Internal server error: Could not decode json"))
		return
	}

	var response datatypes.SearchReponse

	for _, artist := range utils.ArtistsJson {
		// Artist name search
		if utils.Match(s.Search, artist.Name) {
			response.Artists = append(response.Artists, artist)
		}

		// Member search
		for _, member := range artist.Members {
			if utils.Match(s.Search, member) {
				t := "member of "
				if len(artist.Members) == 1 {
					t = "real name of "
				}
				d := datatypes.MemberResponse{
					Name:   member,
					Artist: artist,
					Type:   t,
				}
				response.Members = append(response.Members, d)
			}
		}

		locations := utils.GetLocations(artist.Id)
		for _, location := range locations {
			city := strings.Split(location, "-")[0]
			country := strings.Split(location, "-")[1]
			if utils.Match(city, s.Search) || utils.Match(country, s.Search) {
				response.Locations = append(response.Locations, datatypes.LocationResponse{
					Name:     location,
					Artist:   artist.Name,
					ArtistId: artist.Id,
				})
			}
		}

		// First album date search
		if strings.Contains(artist.FirstAlbum, s.Search) {
			response.FirstAlbumDates = append(response.FirstAlbumDates, datatypes.DateResponse{
				Date:     artist.FirstAlbum,
				Artist:   artist.Name,
				ArtistId: artist.Id,
			})
		}

		// Creation date search
		if strconv.Itoa(artist.CreationDate) == s.Search {
			response.CreationDates = append(response.CreationDates, datatypes.DateResponse{
				Date:     strconv.Itoa(artist.CreationDate),
				Artist:   artist.Name,
				ArtistId: artist.Id,
			})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
