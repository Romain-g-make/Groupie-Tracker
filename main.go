package main

import (
	"fonction/Fonction_go"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var artistes []fonction_go.Artist

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	request_art()
	p := r.URL.Query().Get("pagination")
	pagination_act, _ := strconv.Atoi(p)

	pa := r.URL.Query().Get("page")
	page_act, _ := strconv.Atoi(pa)
	
	tmpl, err := template.ParseFiles("static/page_connexion.html")
	
	var artiste_dec [][]fonction_go.Artist

	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		log.Println("Erreur template :", err)
		return
	}
	if pagination_act != 0 {
		artiste_dec = fonction_go.Pagination(pagination_act,artistes)
	} else{
		artiste_dec = fonction_go.Pagination(1,artistes)
	}

	donnees := fonction_go.Donnees{
		Artist:     artiste_dec[page_act],
		Page:       page_act,
		Pagination: pagination_act,
	}
	fmt.Println(donnees)
	fmt.Println(artistes)

	err = tmpl.Execute(w, donnees)

	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		log.Println("Erreur template :", err)
		return
	}
}

func request_art() {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &artistes)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	handler := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", handler)
	http.HandleFunc("/", renderTemplate)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
