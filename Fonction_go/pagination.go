package fonction_go

func Pagination(nb int, artistes []Artist) [][]Artist {
	var tab_final [][]Artist
	for i := 0; i < nb; i++ {
		tab_final = append(tab_final, artistes[i:(len(artistes)/nb)+i*(len(artistes)/nb)])
	}
	return tab_final
}
