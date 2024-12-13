package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetLocation -
func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	fmt.Printf("Debug: Making request to URL: %s\n", url)

	if val, ok := c.cache.Get(url); ok {
		fmt.Println("Debug: Found in cache")
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	fmt.Println("Debug: Not in cache, making HTTP request")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Debug: Error creating request: %v\n", err)
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Debug: Error making request: %v\n", err)
		return Location{}, err
	}
	defer resp.Body.Close()

	fmt.Printf("Debug: Got response with status: %s\n", resp.Status)

	fmt.Println("Debug: About to read response body")
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Debug: Error reading body: %v\n", err)
		return Location{}, err
	}
	fmt.Printf("Debug: Read %d bytes from response\n", len(dat))

	fmt.Println("Debug: About to unmarshal JSON")
	locationResp := Location{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		fmt.Printf("Debug: Error unmarshaling JSON: %v\n", err)
		return Location{}, err
	}
	fmt.Printf("Debug: Location struct: %+v\n", locationResp)
	fmt.Println("Debug: Successfully unmarshaled JSON")
	c.cache.Add(url, dat)

	return locationResp, nil
}
