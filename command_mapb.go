// command to display the name of next 20 location areas

package main

import (
	"errors"
	"fmt"
	"net/url"
)

func commandMapb() error {
	fullURL := ""

	if conf.Previous != nil {
		prevUrl, err := url.Parse(*conf.Previous)
		if err != nil {
			return err
		}
		query := prevUrl.RawQuery
		fullURL = BASE_LOCATION_URL + "?" + query
	} else {
		return errors.New("On the First Page")
	}

	areas, err := pokeApi(fullURL)
	if err != nil {
		return err
	}

	for _, areaName := range areas.Results {
		fmt.Println(areaName.Name)
	}

	conf.Next = areas.Next
	conf.Previous = areas.Previous

	return nil
}
