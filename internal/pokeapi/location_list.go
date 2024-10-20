// Make the Http request to the PokeAPI.
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Vikuuu/pokedex/internal/pokecache"
)

func LocPokeApi(
	url *string,
	cacheP *pokecache.Cache,
) (areaNames, error) {
	fullURL := baseURL + "/location-area"
	if url != nil {
		fullURL = *url
	}
	var areas areaNames

	if _, found := cacheP.Get(fullURL); !found {
		res, err := http.Get(fullURL)
		if err != nil {
			return areaNames{}, err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return areaNames{}, err
		}

		// Added the result in Cache
		cacheP.Add(fullURL, data)

		err = json.Unmarshal(data, &areas)
		if err != nil {
			return areaNames{}, err
		}
		return areas, nil
	}

	data, _ := cacheP.Get(fullURL)
	err := json.Unmarshal(data, &areas)
	if err != nil {
		return areaNames{}, err
	}

	return areas, nil
}
