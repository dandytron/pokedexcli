package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		locationAreasResp := RespShallowLocations{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationAreasResp, nil

	}

	//cache miss, continue logic below
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	// Redirects should be handled internally by the go client, so can check above the 300s
	if resp.StatusCode > 399 {
		return RespShallowLocations{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Unmarshall the json data into an instance of our struct
	locationAreasResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreasResp, nil

}
