package pokeapi

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
)

func random(exp int) bool {
	catchChance := 100 - (exp / 3)
	randNum := rand.Intn(100)
	if catchChance < 1 {
		catchChance = 1
	}

	if randNum >= catchChance {
		return true
	}
	return false
}

func (c *Client) CatchingPokemon(name string) (pokemonCaught bool, pokemonResp Pokemon, err error) {
	url := baseURL + "/pokemon/" + name

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, Pokemon{}, err
	}

	pokemonData := Pokemon{}
	err = json.Unmarshal(data, &pokemonData)
	if err != nil {
		return false, Pokemon{}, err
	}

	caught := random(pokemonData.BaseExperience)
	return caught, pokemonData, nil
}
