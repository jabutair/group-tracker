package models

// Artist basic info
type Artist struct {
	ID         int      `json:"id"`
	Image      string   `json:"image"`
	Name       string   `json:"name"`
	Members    []string `json:"members"`
	Creation   int      `json:"creationDate"`
	FirstAlbum string   `json:"firstAlbum"`
}

// Locations list
type Location struct {
	Locations []string `json:"locations"`
}

// All locations response
type AllLocation struct {
	Index []struct {
		Locations []string `json:"locations"`
	} `json:"index"`
}

// Concert dates
type Dates struct {
	Dates []string `json:"dates"`
}

// Dates with locations
type Relationship struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Full artist details page data
type Details struct {
	Artist            Artist
	Locations         Location
	ConcertDates      Dates
	DatesAndLocations Relationship
}
