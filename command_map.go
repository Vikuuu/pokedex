// command to display the name of next 20 location areas

package main

import (
	"fmt"
	"net/url"
)

func commandMap() error {
	fullURL := BASE_LOCATION_URL

	if conf.Next != nil {
		nextUrl, err := url.Parse(*conf.Next)
		if err != nil {
			return err
		}
		query := nextUrl.RawQuery
		fullURL += "?" + query
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
