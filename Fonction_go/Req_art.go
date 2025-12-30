package fonction_go

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Request_art(artistes *[]Artist) error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return fmt.Errorf("appel API artists: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API artists retourne le statut %d", resp.StatusCode)
	}

	var payload []Artist
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return fmt.Errorf("décodage JSON artists: %w", err)
	}

	converted := make([]Artist, 0, len(payload))
	for _, a := range payload {
		converted = append(converted, Artist{
			Id:            a.Id,
			Nom:           a.Nom,
			Image:         a.Image,
			Annee_deb:     a.Annee_deb,
			Date_prem_alb: a.Date_prem_alb,
			Membres:       a.Membres,
		})
	}

	*artistes = converted
	return nil
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
