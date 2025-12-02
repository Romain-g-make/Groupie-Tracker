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

func Get_artistes(artistes []Artist) []string {
	var list_art []string
	for i:=0 ; i<len(artistes); i++ {
		list_art = append(list_art, artistes[i].Nom)
	}
	return list_art
}

func Find_artist(artistes []Artist, name string) Artist {
	for i:=0 ; i<len(artistes); i++ {
		if artistes[i].Nom == name {
			return artistes[i]
		}
	}
	return Artist{}
}