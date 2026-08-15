package handlers

import (
	"html/template"
	"net/http"

	"group-tracker/models"
	"group-tracker/services"
)

type PageData struct {
	Artists      []models.Artist
	MembersRange []int
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))

	var artists []models.Artist
	var locations models.AllLocation

	services.FetchJSON("https://groupietrackers.herokuapp.com/api/artists", &artists)
	services.FetchJSON("https://groupietrackers.herokuapp.com/api/locations", &locations)

	filtered := services.FilterArtists(
		artists,
		locations,
		r.URL.Query().Get("creationDateStart"),
		r.URL.Query().Get("creationDateEnd"),
		r.URL.Query().Get("firstAlbumStart"),
		r.URL.Query().Get("firstAlbumEnd"),
		r.URL.Query()["members"],
		r.URL.Query().Get("location"),
		r.URL.Query().Get("search"),
	)

	data := PageData{
		Artists:      filtered,
		MembersRange: []int{1, 2, 3, 4, 5, 6, 7, 8},
	}

	tmpl.Execute(w, data)
}
