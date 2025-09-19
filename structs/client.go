package structs

import (
	"net/http"
	"time"
	"io"
	"encoding/json"
)

// Client -
type Client struct {
	HttpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client {
		HttpClient: http.Client {
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

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	res, err := client.HttpClient.Do(req)
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

	return locations, nil
}