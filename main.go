package main

import (
	fonction_go "fonction/Fonction_go"
	"io"
	"log"
	"net/http"
)

func main() {

	// Simple proxy to avoid CORS issues when the pages fetch the external API.
	// The browser will call /api/artists on the same origin and Go will forward the request.
	http.HandleFunc("/api/artists", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			http.Error(w, "Erreur proxy vers API distante", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})

	// http.HandleFunc("/concerts", fonction_go.ConcertsHandler)

	handler := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", handler)
	http.HandleFunc("/", fonction_go.RenderTemplate)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
