package datatypes

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type SearchRequest struct {
	Search string
}

type SearchReponse struct {
	Artists         []Artist           `json:"artists"`
	Members         []MemberResponse   `json:"members"`
	Locations       []LocationResponse `json:"locations"`
	FirstAlbumDates []DateResponse     `json:"firstAlbumDates"`
	CreationDates   []DateResponse     `json:"creationDates"`
}

type LocationResponse struct {
	Name     string `json:"name"`
	Artist   string `json:"artist"`
	ArtistId int    `json:"artistId"`
}

type DateResponse struct {
	Date     string `json:"date"`
	Artist   string `json:"artist"`
	ArtistId int    `json:"artistId"`
}

type MemberResponse struct {
	Name   string `json:"name"`
	Artist Artist `json:"artist"`
	Type   string `json:"type"`
}

type ArtistResponse struct {
	Artist    Artist              `json:"artist"`
	Relations Relation            `json:"relations"`
	Locations map[string][]string `json:"locations"`
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type ArtistsPage struct {
	Artists   []Artist
	Locations []string
}

type FiltersRequest struct {
	FromCreationYear   int    `json:"fromCreationYear"`
	ToCreationYear     int    `json:"toCreationYear"`
	FromFirstAlbumYear int    `json:"fromAlbumYear"`
	ToFirstAlbumYear   int    `json:"toAlbumYear"`
	MembersMin         int    `json:"membersMin"`
	MembersMax         int    `json:"membersMax"`
	Location           string `json:"location"`
}

type ErrorPage struct {
	ErrStatus string
	ErrMsg    string
}
