package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) ListLocations(pageUrl *string) (LocationAreaList, error) {
	url := baseURL + "/location-area"

	// If a pageUrl is provided, use it instead
	if pageUrl != nil {
		url = *pageUrl
	}

	// Create the HTTP request
	request, err := http.NewRequest("GET", url, nil)

	// Check for errors while creating the request
	if err != nil {
		fmt.Println("1")
		return LocationAreaList{}, err
	}

	// Execute the HTTP request
	response, err := client.httpClient.Do(request)

	// Check for errors while executing the request
	if err != nil {
		fmt.Println("2")
		return LocationAreaList{}, err
	}

	// Defer the closing of the response body
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)

	// Check for errors while reading the response body
	if err != nil {
		fmt.Println("3")
		return LocationAreaList{}, err
	}

	// Declare a variable to hold the location area data
	locationArea := LocationAreaList{}

	// Unmarshal the response body into the location area struct
	err = json.Unmarshal(body, &locationArea)

	// Check for errors while unmarshaling the response body
	if err != nil {
		fmt.Println("4")
		return LocationAreaList{}, err
	}

	// Return the location area data
	return locationArea, nil
}
