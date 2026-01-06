package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(configuration *config) error {
	res, err := http.Get(configuration.nextURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var result locationAreaResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return err
	}

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	configuration.nextURL = result.Next
	configuration.previousURL = result.Previous
	return nil
}
