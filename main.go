package main

import (
	fonction_go "fonction/Fonction_go"
	"log"
	"net/http"
)

func main() {

	handler := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", handler)
	http.HandleFunc("/", fonction_go.RenderTemplate)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
