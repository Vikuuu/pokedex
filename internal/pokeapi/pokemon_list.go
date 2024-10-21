package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonInArea(area string) (pokemon, error) {
	url := baseURL + "/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return pokemon{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokemon{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemon{}, err
	}

	pokemonResp := pokemon{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return pokemon{}, err
	}

	c.cache.Add(url, data)
	return pokemonResp, nil
}
