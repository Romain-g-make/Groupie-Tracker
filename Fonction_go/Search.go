package fonction_go

import (
	"fmt"
	"strings"
)

func Search(artistes []Artist, find_art string) map[string]int {

	dic_art := make(map[string]int)
	find_art_lower := strings.ToLower(find_art)

	for i := 0; i < len(artistes); i++ {
		correspondance := 0
		nom_lower := strings.ToLower(artistes[i].Nom)
		if nom_lower == find_art_lower {
			correspondance = len(find_art) * 3
		} else if strings.Contains(nom_lower, find_art_lower) {
			correspondance = len(find_art) * 2
			if nom_lower[0:1] == find_art_lower[0:1] {
				correspondance++
			}
		} else {
			min_len := len(nom_lower)
			if len(find_art_lower) < min_len {
				min_len = len(find_art_lower)
			}
			for j := 0; j < min_len; j++ {
				if nom_lower[j] == find_art_lower[j] {
					correspondance++
				} else {
					break
				}
			}
		}
		if correspondance > 0 {
			dic_art[artistes[i].Nom] = correspondance
		}
	}
	return dic_art
}

func Searching(artistes []Artist, find_art string) []Artist {

	var result []Artist

	if find_art == "" {
		return artistes
	}
	dic_art := Search(artistes, find_art)
	if len(dic_art) == 0 {
		fmt.Println("Aucun résultat trouvé pour:", find_art)
		return result
	}
	for len(dic_art) > 0 {
		max_correspondance := 0
		max_key := ""
		for key, score := range dic_art {
			if score > max_correspondance {
				max_correspondance = score
				max_key = key
			}
		}
		if max_key != "" {
			artiste := Find_artist(artistes, max_key)
			if artiste.Nom != "" {
				result = append(result, artiste)
			}
			delete(dic_art, max_key)
		}
	}
	return result
}
