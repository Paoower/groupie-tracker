package utils

import (
	"groupie-tracker/datatypes"
)

func GetAllLocations(json []datatypes.Location) []string {
	locs := []string{}
	for _, e := range json {
		locs = append(locs, e.Locations...)
	}
	return locs
}

func GetArtistLocations(id int, locations []datatypes.Location, relations []datatypes.Relation) map[string][]string {
	locationDates := make(map[string][]string)

	for _, relation := range relations {
		if relation.Id == id {
			for location, dates := range relation.DatesLocations {
				locationDates[location] = dates
			}
			break
		}
	}

	return locationDates
}

func GetLocations(id int) []string {
	response := []string{}
	for _, locations := range LocationsJson {
		if locations.Id == id {
			response = locations.Locations
		}
	}
	return response
}
