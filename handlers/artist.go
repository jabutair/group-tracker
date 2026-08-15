package handlers

import (
	"html/template"
	"net/http"

	"group-tracker/models"
	"group-tracker/services"
)

var details models.Details

func Artist(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("artistNumber")

	services.FetchJSON("https://groupietrackers.herokuapp.com/api/artists/"+id, &details.Artist)
	services.FetchJSON("https://groupietrackers.herokuapp.com/api/locations/"+id, &details.Locations)
	services.FetchJSON("https://groupietrackers.herokuapp.com/api/dates/"+id, &details.ConcertDates)
	services.FetchJSON("https://groupietrackers.herokuapp.com/api/relation/"+id, &details.DatesAndLocations)

	tmpl := template.Must(template.ParseFiles("templates/artist.html"))
	tmpl.Execute(w, details)
}
