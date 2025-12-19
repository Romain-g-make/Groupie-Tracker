package fonction_go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


func Request_art(artistes *[]Artist) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &artistes)
	if err != nil {
		fmt.Println(err)
	}
}
