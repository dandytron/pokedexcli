package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		fmt.Println("Cache hit.")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreasResp, nil

	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	// Redirects should be handled internally by the go client, so can check above the 300s
	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// Unmarshall the json data into an instance of our struct
	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreasResp, nil

}
