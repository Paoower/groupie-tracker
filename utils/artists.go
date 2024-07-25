package utils

import (
	"groupie-tracker/datatypes"
	"strings"
)

func GetArtist(id int, artistsJson []datatypes.Artist) datatypes.Artist {
	a := datatypes.Artist{}
	for _, artist := range artistsJson {
		if artist.Id == id {
			a = artist
		}
	}
	return a
}

func GetRelations(id int, relationsJson []datatypes.Relation) datatypes.Relation {
	r := datatypes.Relation{}
	for _, relation := range relationsJson {
		if relation.Id == id {
			r = relation
		}
	}
	return r
}

func GetLocation(id int, locations []datatypes.Location) []string {
	l := []string{}
	for _, location := range locations {
		if location.Id == id {
			l = location.Locations
		}
	}

	for i := 0; i < len(l); i++ {
		l[i] = strings.ReplaceAll(l[i], "-", " ")
		l[i] = strings.ReplaceAll(l[i], "_", " ")
		l[i] = strings.Title(l[i])
	}

	return l
}
