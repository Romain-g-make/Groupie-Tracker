package fonction_go

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var artistes []Artist

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	Request_art(&artistes)
	p := r.URL.Query().Get("pagination")
	pagination_act, _ := strconv.Atoi(p)

	pa := r.URL.Query().Get("page")
	page_act, _ := strconv.Atoi(pa)

	is_art := r.URL.Query().Get("Is_artist")

	tmpl, err := template.ParseFiles("static/page_connexion.html")
	var current_artist Artist

	if is_art != "" {
		current_artist = Find_artist(artistes, is_art)
		tmpl, err = template.ParseFiles("static/artist_detail.html")
	} else {
		current_artist = Artist{}
	}

	if r.URL.Query().Get("Search") != "" {
		artistes = Searching(artistes, r.URL.Query().Get("Search"))
		pagination_act = 1
		page_act = 0
		tmpl, err = template.ParseFiles("static/artists_list.html")
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

	alphabet := Get_alphabet(artistes)

	donnees := Donnees{
		Artist:     artistes,
		Page:       page_act,
		Pagination: pagination_act,
		Lettre:     alphabet,
		Is_artist:  current_artist,
	}

	err = tmpl.Execute(w, donnees)

	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		log.Println("Erreur template :", err)
		return
	}
}
