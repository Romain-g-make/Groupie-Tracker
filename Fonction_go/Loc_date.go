package fonction_go

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Get_Locations(id int) []string {

	var locations []string
	url := "https://groupietrackers.herokuapp.com/api/locations" + "/" + string(id)
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &locations)
	return locations
}

func Get_Dates(id int) []string {

	var dates []string
	url := "https://groupietrackers.herokuapp.com/api/dates" + "/" + string(id)
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &dates)
	return dates
}