package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationData(pageUrl *string) Location {
	if pageUrl == nil {
		defaultUrl := "https://pokeapi.co/api/v2/location-area/"
		pageUrl = &defaultUrl
	}
	res, err := http.Get(*pageUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s", body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var locationData Location
	err = json.Unmarshal(body, &locationData)
	if err != nil {
		log.Fatal(err)
	}

	return locationData
}
