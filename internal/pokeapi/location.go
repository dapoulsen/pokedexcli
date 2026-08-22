package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dapoulsen/pokedexcli/internal/pokecache"
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

const baseURL = "https://pokeapi.co/api/v2/"

func GetLocationData(pageUrl *string, cache *pokecache.Cache) (Location, error) {
	if pageUrl == nil {
		defaultUrl := baseURL + "location-area/"
		pageUrl = &defaultUrl
	}

	e, ok := cache.Get(*pageUrl)
	if ok {
		var locationData Location
		err := json.Unmarshal(e, &locationData)
		if err != nil {
			return Location{}, err
		}
		return locationData, nil
	}

	res, err := http.Get(*pageUrl)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		return Location{}, fmt.Errorf("Error: %s", body)
	}
	if err != nil {
		return Location{}, err
	}

	var locationData Location
	err = json.Unmarshal(body, &locationData)
	if err != nil {
		return Location{}, err
	}

	cache.Add(*pageUrl, body)

	return locationData, nil
}
