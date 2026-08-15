package services

import (
	"strconv"
	"strings"

	"group-tracker/models"
)

// Filters artists based on all query params
func FilterArtists(
	artists []models.Artist,
	locs models.AllLocation,
	cdS, cdE, faS, faE string,
	members []string,
	location, search string,
) []models.Artist {

	var result []models.Artist
	search = strings.ToLower(search)

	if len(artists) > 20 {
		artists[20].Image = "https://brasilia.deboa.com/wp-content/uploads/2023/12/Mamonas-Assassinas-O-Filme.jpg"
	}

	for _, artist := range artists {
		match := true

		// Search by name or members
		if search != "" {
			found := strings.Contains(strings.ToLower(artist.Name), search)
			for _, m := range artist.Members {

				if strings.Contains(strings.ToLower(m), search) {
					found = true
				}
			}
			if !found {
				match = false
			}
		}

		// Creation date filter
		if match && cdS != "" && cdE != "" {
			s, _ := strconv.Atoi(cdS)
			e, _ := strconv.Atoi(cdE)
			if artist.Creation < s || artist.Creation > e {
				match = false
			}
		}

		// First album filter
		if match && faS != "" && faE != "" {
			s, _ := strconv.Atoi(faS)
			e, _ := strconv.Atoi(faE)
			year, _ := strconv.Atoi(artist.FirstAlbum[len(artist.FirstAlbum)-4:])
			if year < s || year > e {
				match = false
			}
		}

		// Members count filter
		if match && len(members) > 0 {
			ok := false
			for _, m := range members {
				n, _ := strconv.Atoi(m)
				if len(artist.Members) == n {
					ok = true
				}
			}
			if !ok {
				match = false
			}
		}

		// Location filter
		if match && location != "" {
			found := false
			for _, l := range locs.Index[artist.ID-1].Locations {
				if strings.Contains(strings.ToLower(l), strings.ToLower(location)) {
					found = true
				}
			}
			if !found {
				match = false
			}
		}

		if match {
			result = append(result, artist)
		}
	}

	return result
}
