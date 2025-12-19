package fonction_go

func Search(artistes []Artist, find_art string) (Artist, int) {
	result := Artist{}
	var correspondance int
	for i := 0; i < len(artistes); i++ {
		if artistes[i].Nom == find_art {
			result = artistes[i]
			correspondance = len(find_art)
		} else {
			min_len := len(artistes[i].Nom)
			if len(find_art) < min_len {
				min_len = len(find_art)
			}
			for j := 0; j < min_len; j++ {
				if artistes[i].Nom[j] == find_art[j] {
					correspondance++
				} else {
					break
				}
			}
		}
	}
	return result, correspondance
}

func Searching(artistes []Artist, find_art string) []Artist {
	var result []Artist
	var dic_art map[string]int
	for i := 0; i < 2; i++ {
		i--
		artiste, correspondance := Search(artistes, find_art)
		dic_art[artiste.Nom] = correspondance
		find_art = find_art[:len(find_art)-1]
		if len(find_art) == 0 {
			break
		}
	}
	for key := range dic_art {
		max_correspondance := dic_art[key]
		for key2 := range dic_art {
			if dic_art[key2] > max_correspondance {
				max_correspondance = dic_art[key2]
				key = key2
			}
		}
		artiste := Find_artist(artistes, key)
		result = append(result, artiste)
	}
	return result
}
