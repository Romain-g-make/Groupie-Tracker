package fonction_go

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Request_art() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, fmt.Errorf("appel API artists: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API artists retourne le statut %d", resp.StatusCode)
	}

	var payload []Artist
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("décodage JSON artists: %w", err)
	}

	artistes := make([]Artist, 0, len(payload))
	for _, a := range payload {
		artistes = append(artistes, Artist{
			Id:            a.Id,
			Nom:           a.Nom,
			Image:         a.Image,
			Annee_deb:     a.Annee_deb,
			Date_prem_alb: a.Date_prem_alb,
			Membres:       a.Membres,
		})
	}

	return artistes, nil
}

func Get_artistes(artistes []Artist) []string {
	var list_art []string
	for i := 0; i < len(artistes); i++ {
		list_art = append(list_art, artistes[i].Nom)
	}
	return list_art
}

func Find_artist(artistes []Artist, name string) Artist {
	for i := 0; i < len(artistes); i++ {
		if artistes[i].Nom == name {
			return artistes[i]
		}
	}
	return Artist{}
}
