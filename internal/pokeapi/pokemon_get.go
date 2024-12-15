package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetLocation -
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	// Check if there's a cache hit; get it from the cache and return it
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	// Pokemon data not found in cache already, construct HTTP request for pokemon json
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// Actually make the request. This block executes the request using the client's HTTP properties.
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	// The response body is read into a byte slice using io.ReadAll
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// Unmarshal JSON
	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	// Add the resp to the cache so the next time this pokemon is requested we can skip the call
	c.cache.Add(url, dat)

	return pokemonResp, nil
}
