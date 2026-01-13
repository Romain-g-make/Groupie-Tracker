package fonction_go

import (
	"fmt"
	"strings"
)

func getSearchField(artist Artist, searchType string) []string {
	switch searchType {
	case "artist":
		return []string{artist.Nom}
	case "date":
		return Get_Dates(artist.Id)
	case "location":
		return Get_Locations(artist.Id)
	default:
		return []string{}
	}
}

func Search(artistes []Artist, find string, searchType string) map[string]int {

	dic_art := make(map[string]int)
	findLower := strings.ToLower(find)

	for _, artiste := range artistes {

		field := strings.ToLower(strings.Join(getSearchField(artiste, searchType), " "))

		correspondance := 0

		if field == findLower {
			correspondance = len(find) * 3
		} else if strings.Contains(field, findLower) {
			correspondance = len(find) * 2
			if field[0:1] == findLower[0:1] {
				correspondance++
			}
		} else {
			minLen := len(field)
			if len(findLower) < minLen {
				minLen = len(findLower)
			}
			for i := 0; i < minLen; i++ {
				if field[i] == findLower[i] {
					correspondance++
				} else {
					break
				}
			}
		}

		if correspondance > 0 {
			dic_art[artiste.Nom] = correspondance
		}
	}
	return dic_art
}

func Searching(artistes []Artist, find string, searchType string) []Artist {

	var result []Artist

	if find == "" {
		return artistes
	}

	dic_art := Search(artistes, find, searchType)

	if len(dic_art) == 0 {
		fmt.Println("Aucun résultat trouvé pour:", find)
		return result
	}

	for len(dic_art) > 0 {
		maxScore := 0
		maxKey := ""

		// Parcourir dans l'ordre original pour avoir un tri déterministe
		for _, artiste := range artistes {
			if score, exists := dic_art[artiste.Nom]; exists && score > maxScore {
				maxScore = score
				maxKey = artiste.Nom
			}
		}

		if maxKey != "" {
			artiste := Find_artist(artistes, maxKey)
			if artiste.Nom != "" {
				result = append(result, artiste)
			}
			delete(dic_art, maxKey)
		} else {
			break
		}
	}

	return result
}
