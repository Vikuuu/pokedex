// Make the Http request to the PokeAPI.
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocPokeApi(url *string) (areaNames, error) {
	fullURL := baseURL + "/location-area"
	if url != nil {
		fullURL = *url
	}

	if val, ok := c.cache.Get(fullURL); ok {
		locationResp := areaNames{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return areaNames{}, nil
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return areaNames{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return areaNames{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return areaNames{}, err
	}

	locationResp := areaNames{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return areaNames{}, err
	}

	c.cache.Add(fullURL, data)
	return locationResp, nil
}
