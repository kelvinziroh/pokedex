package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap() error {
	data := getLocationAreas()
	locationAreas := data.Results
	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}
	return nil
}

func getLocationAreas() LocationAreas {
	res, httpErr := http.Get("https://pokeapi.co/api/v2/location-area/")
	if httpErr != nil {
		fmt.Printf("HTTP Error: %v", httpErr)
	}
	defer res.Body.Close()
	var locationAreas LocationAreas
	parseErr := json.NewDecoder(res.Body).Decode(&locationAreas)
	if parseErr != nil {
		fmt.Printf("Parse Error: %v", parseErr)
	}
	return locationAreas
}

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
