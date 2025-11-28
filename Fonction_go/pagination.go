package Fonction_go

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

func Pagination(nb int, artistes Artist) []Artist {
	var tab_final []Artist
	for i := 0; i < nb; i++ {
		tab_final = append(tab_final, artistes[i:(len(artistes)/nb)+i*(len(artistes)/nb)])
	}
	return tab_final
}

//Convertir la structure artiste en tableau d'artiste car pour l'instant on a : Tous les noms d'un côté, toutes les dates d'un autre ect..
//alors qu'on veut : un artiste avec ses dates ect... puis un autre artiste ...
