package main

type config struct {
	nextURL     string
	previousURL string
}

type cliCommand struct {
	name        string
	description string
	callback    func(configuration *config) error
}

type locationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
