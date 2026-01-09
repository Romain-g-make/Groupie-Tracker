package fonction_go

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var artistes []Artist

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	current_art_loc_dates := map[string]string{}
	Reset := r.URL.Query().Get("Reset")
	is_art := r.URL.Query().Get("Is_artist")

	if Reset != "list" {
		if err := Request_art(&artistes); err != nil {
			http.Error(w, "Erreur lors de la récupération des artistes", http.StatusBadGateway)
			log.Println("Erreur Request_art :", err)
			return
		}
	}

	p := r.URL.Query().Get("pagination")
	pagination_act, _ := strconv.Atoi(p)

	pa := r.URL.Query().Get("page")
	page_act, _ := strconv.Atoi(pa)

	tmpl, err := template.ParseFiles("static/page_connexion.html")
	var current_artist Artist

	if r.URL.Query().Get("Search") != "" || Reset == "list" {
		artistes = Searching(artistes, r.URL.Query().Get("Search"))
		pagination_act = 0
		page_act = 0
		tmpl, err = template.ParseFiles("static/artists_list.html")
	}

	if is_art != "" {
		current_artist = Find_artist(artistes, is_art)
		current_art_loc_dates = Loc_date(current_artist.Id)
		tmpl, err = template.ParseFiles("static/artist_detail.html")
	} else {
		current_artist = Artist{}
	}

	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		log.Println("Erreur template :", err)
		return

	}
	if pagination_act != 0 {
		artiste_dec := Pagination(pagination_act, artistes)
		artistes = artiste_dec[page_act]
	}

	lettres := Get_alphabet(artistes)
	searchTerm := r.URL.Query().Get("Search")

	donnees := Donnees{
		Artist:     artistes,
		Page:       page_act,
		Pagination: pagination_act,
		Lettres:    lettres,
		Is_artist:  current_artist,
		Loc_dates:  current_art_loc_dates,
		Search:     searchTerm,
	}

	err = tmpl.Execute(w, donnees)

	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		log.Println("Erreur template :", err)
		return
	}
}
