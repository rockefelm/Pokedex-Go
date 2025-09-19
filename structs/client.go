package structs

import (
	"net/http"
	"time"
	"io"
	"encoding/json"
)

// Client -
type Client struct {
	cache		Cache
	httpClient 	http.Client
}

func NewClient(timeout, cacheDelay time.Duration) Client {
	return Client {
		cache: NewCache(cacheDelay),
		httpClient: http.Client {
			Timeout: timeout,
		},
	}
}

const (
	BaseURL = "https://pokeapi.co/api/v2/"
)

func (client *Client) Locations(locationURL *string) (Locations, error) {
	url := BaseURL + "/location-area"
	if locationURL != nil {
		url = *locationURL
	}

	if val, ok := client.cache.Get(url); ok {
		locations := Locations{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return Locations{}, err
		}

		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}

	locations := Locations{}
	err = json.Unmarshal(jsonData, &locations)
	if err != nil {
		return Locations{}, err
	}

	client.cache.Add(url, jsonData)
	return locations, nil
}