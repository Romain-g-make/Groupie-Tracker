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

	var artiste_dec [][]Artist
	var current_artist Artist

	if is_art != "" {
		current_artist = Find_artist(artistes, is_art)
	} else {
		current_artist = Artist{}
	}

	if r.URL.Query().Get("Search") != "" {
		artiste_dec = append(artiste_dec, Searching(artistes, r.URL.Query().Get("Search")))
	}

	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		log.Println("Erreur template :", err)
		return
	}
	if pagination_act != 0 {
		artiste_dec = Pagination(pagination_act, artistes)
	} else {
		artiste_dec = Pagination(1, artistes)
	}

	alphabet := Get_alphabet(artistes)

	donnees := Donnees{
		Artist:     artiste_dec[page_act],
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
