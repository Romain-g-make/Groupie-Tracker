package fonction_go

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Get_Locations(id int) []string {
	type locationsResp struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
		DatesLink string   `json:"dates"`
	}

	url := "https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil
	}

	var payload locationsResp
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		return nil
	}
	return payload.Locations
}

func Get_Dates(id int) []string {
	type datesResp struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	}

	url := "https://groupietrackers.herokuapp.com/api/dates/" + strconv.Itoa(id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil
	}

	var payload datesResp
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		return nil
	}
	return payload.Dates
}

func Loc_date(id int) (map[string]string) {
	locations := Get_Locations(id)
	dates := Get_Dates(id)
	result := make(map[string]string)
	for i:=0 ; i < len(locations) && i < len(dates); i++ {
		result[locations[i]] = dates[i]
	}
	return result
}