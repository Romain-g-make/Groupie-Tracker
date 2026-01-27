package fonction_go

import "sort"

func Tri_alpha_croissant(artistes []Artist) []Artist {
	result := make([]Artist, len(artistes))
	copy(result, artistes)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Nom < result[j].Nom
	})
	return result
}

func Tri_alpha_decroissant(artistes []Artist) []Artist {
	result := make([]Artist, len(artistes))
	copy(result, artistes)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Nom > result[j].Nom
	})
	return result
}

func Tri_anciennete_croissant(artistes []Artist) []Artist {
	result := make([]Artist, len(artistes))
	copy(result, artistes)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Annee_deb < result[j].Annee_deb
	})
	return result
}

func Tri_anciennete_decroissant(artistes []Artist) []Artist {
	result := make([]Artist, len(artistes))
	copy(result, artistes)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Annee_deb > result[j].Annee_deb
	})
	return result
}
