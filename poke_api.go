// Make the Http request to the PokeAPI.
package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func pokeApi(url string) (areaNames, error) {
	res, err := http.Get(url)
	if err != nil {
		return areaNames{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return areaNames{}, err
	}

	var areas areaNames
	err = json.Unmarshal(data, &areas)
	if err != nil {
		return areaNames{}, err
	}
	return areas, nil
}
