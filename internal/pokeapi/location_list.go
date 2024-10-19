// Make the Http request to the PokeAPI.
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func LocPokeApi(url *string) (areaNames, error) {
	fullURL := baseURL + "/location-area"
	if url != nil {
		fullURL = *url
	}

	res, err := http.Get(fullURL)
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
