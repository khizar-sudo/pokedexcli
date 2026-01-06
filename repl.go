package main

import "strings"

func cleanInput(text string) []string {
	var result []string
	text = strings.TrimSpace(text)
	if text == "" {
		return result
	}
	for _, word := range strings.Fields(text) {
		result = append(result, word)
	}
	return result
}