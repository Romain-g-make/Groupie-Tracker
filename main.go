package main

import (
	"fonction/Fonction_go"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Artist struct {
	id            int      `json:"id"`
	nom           string   `json:"name"`
	image         string   `json:"image"`
	annee_deb     int      `json:"creationDate"`
	date_prem_alb string   `json:"firstAlbum"`
	membres       []string `json:"members"`
	locations     []string `json:"locations"`
	date_concerts []string `json:"concertDates"`
}

var artistes Artist

type Donnees struct {
	artist Artist
	page int
	pagination int
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	
	p := r.URL.Query().Get(pagination)
	pagination_act,_ := strconv.Atoi(p)
	
	pa := r.URL.Query().Get(page)
	page_act,_ := strconv.Atoi(pa)
	
	tmpl, err := template.ParseFiles("static/page_connexion")
	artiste_dec := []Artist{artistes}
	
	if pagination_act != 0 {
		artiste_dec = Fonction_go.Pagination(pagination_act,artiste_dec)
		
	}
	
	donnees := Donnees{
		artist : artiste_dec[page_act] ,
		page: page_act,
		pagination: pagination_act,
	}
	
	err = tmpl.Execute(w, donnees)
	
	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		log.Println("Erreur template :", err)
		return
	}
}

func main() {

	handler := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", handler)
	http.HandleFunc("/", renderTemplate)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
