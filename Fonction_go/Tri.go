package fonction_go

func Search (artistes []Artist, find_art string) Artist {
	result := Artist{}
	for i:=0 ; i<len(artistes); i++ {
		if artistes[i].Nom == find_art {
			result = artistes[i]
		}
	}
	return result
}

func Searching (artistes []Artist, find_art string) []Artist{
	var result []Artist
	for i:=0 ; i<2 ; i++{
		i--
		artiste := Search(artistes,find_art)
		if artiste.Nom != "" {
			result = append (result, artiste)
			i++
			find_art = find_art[:len(find_art)-1]
		}
	}
	return result
}

func Tri_alpha_croissant (artistes []Artist) []Artist {
	for i:=0 ; i<len(artistes)-1 ; i++ {
		for j:=i+1 ; j<len(artistes) ; j++ {
			if artistes[i].Nom > artistes[j].Nom {
				temp := artistes[i]
				artistes[i] = artistes[j]
				artistes[j] = temp
			}
		}
	}
	return artistes
}

func Tri_alpha_decroissant (artistes []Artist) []Artist {
	for i:=0 ; i<len(artistes)-1 ; i++ {
		for j:=i+1 ; j<len(artistes) ; j++ {
			if artistes[i].Nom < artistes[j].Nom {
				temp := artistes[i]
				artistes[i] = artistes[j]
				artistes[j] = temp
			}
		}
	}
	return artistes
}