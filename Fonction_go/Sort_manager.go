package fonction_go

func tri_artists(artists []Artist, sortOrder string) []Artist {
	switch sortOrder {
	case "asc":
		return Tri_alpha_croissant(artists)
	case "desc":
		return Tri_alpha_decroissant(artists)
	default:
		return artists
	}
}
func options_tri() map[string]string {
	return map[string]string{
		"":     "Sans tri",
		"asc":  "A → Z (Croissant)",
		"desc": "Z → A (Décroissant)",
	}
}
