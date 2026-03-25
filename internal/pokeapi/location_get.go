package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(areaName string) (Location, error) {
	url := baseURL + "/location-area/" + areaName

	// Check Cache
	if value, exist := c.pokeCache.Get(url); exist {
		locationResp := Location{}
		err := json.Unmarshal(value, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	// Make a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	// Add to Cache
	c.pokeCache.Add(url, dat)

	locationResp := Location{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Location{}, err
	}

	return locationResp, nil
}