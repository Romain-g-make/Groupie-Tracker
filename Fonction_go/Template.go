package fonction_go

import (
	"fmt"
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
	
	tmpl, err := template.ParseFiles("static/page_connexion.html")
	
	var artiste_dec [][]Artist

	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		log.Println("Erreur template :", err)
		return
	}
	if pagination_act != 0 {
		artiste_dec = Pagination(pagination_act,artistes)
	} else{
		artiste_dec = Pagination(1,artistes)
	}

	lettres := Get_alphabet(artistes)

	donnees := Donnees{
		Artist:     artiste_dec[page_act],
		Page:       page_act,
		Pagination: pagination_act,
		Lettre:     lettres,
	}
	fmt.Println(donnees)

	err = tmpl.Execute(w, donnees)

	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		log.Println("Erreur template :", err)
		return
	}
}
