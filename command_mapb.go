package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMapb(configuration *config) error {
	if configuration.previousURL == "" {
		fmt.Println("You're on the first page of locations")
		return nil
	}

	res, err := http.Get(configuration.previousURL)
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
