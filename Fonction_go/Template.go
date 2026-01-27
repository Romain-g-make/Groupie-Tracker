package fonction_go

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
	}

	params := extractQueryParams(r)

	baseArtists, err := Request_art()
	if err != nil {
		handleError(w, "Erreur lors de la récupération des artistes", http.StatusBadGateway, err)
		return
	}

	artistesToDisplay := make([]Artist, len(baseArtists))
	copy(artistesToDisplay, baseArtists)

	tmplPath, data := prepareTemplateData(artistesToDisplay, params, funcMap)

	tmpl, err := template.New(tmplPath).Funcs(funcMap).ParseFiles("static/" + tmplPath)
	if err != nil {
		handleError(w, "Erreur lors du chargement du template", http.StatusInternalServerError, err)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		handleError(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError, err)
		return
	}
}

func handleError(w http.ResponseWriter, message string, statusCode int, err error) {
	http.Error(w, message, statusCode)
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
}
