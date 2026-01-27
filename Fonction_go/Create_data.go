package fonction_go

import (
	"html/template"
)

func prepareTemplateData(artists []Artist, params queryParams, funcMap template.FuncMap) (string, Donnees) {
	var tmplPath string
	var currentArtist Artist
	var locDates map[string]string
	displayedArtists := artists

	if params.artistName == "" && (params.searchQuery != "" || params.reset == "list") {
		displayedArtists = Searching(displayedArtists, params.searchQuery, params.searchType)
	}
	
	// Appliquer le tri
	if params.sortOrder == "asc" {
		displayedArtists = Tri_alpha_croissant(displayedArtists)
	} else if params.sortOrder == "desc" {
		displayedArtists = Tri_alpha_decroissant(displayedArtists)
	}
	
	if params.artistName != "" {
		tmplPath = "artist_detail.html"
		currentArtist = Find_artist(artists, params.artistName)
		if currentArtist.Nom != "" {
			locDates = Loc_date(currentArtist.Id)
		}
	} else if params.searchQuery != "" || params.reset == "list" {
		tmplPath = "artists_list.html"

		if params.pagination > 0 {
			totalPages := GetTotalPages(params.pagination, displayedArtists)
			if params.page >= totalPages {
				params.page = 0
			}
			displayedArtists = Pagination(params.pagination, displayedArtists, params.page)
		}
	} else {
		tmplPath = "page_connexion.html"
	}

	data := Donnees{
		Artist:     displayedArtists,
		Page:       params.page,
		Pagination: params.pagination,
		Lettres:    Get_alphabet(artists),
		Is_artist:  currentArtist,
		Loc_dates:  locDates,
		Search:     params.searchQuery,
		SearchType: params.searchType,
		Sort:       params.sortOrder,
	}

	return tmplPath, data
}
