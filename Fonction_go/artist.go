package fonction_go

type Artist struct {
	Id            int      `json:"id"`
	Nom           string   `json:"name"`
	Image         string   `json:"image"`
	Annee_deb     int      `json:"creationDate"`
	Date_prem_alb string   `json:"firstAlbum"`
	Membres       []string `json:"members"`
	Locations     string `json:"locations"`
	Date_concerts string `json:"concertDates"`
}
