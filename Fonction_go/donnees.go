package fonction_go

type Donnees struct {
	Artist     []Artist
	Page       int
	Pagination int
	Lettres    []string
	Is_artist  Artist
	Loc_dates  map[string]string
	Search     string
	SearchType string
}
