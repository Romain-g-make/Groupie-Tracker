package fonction_go

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