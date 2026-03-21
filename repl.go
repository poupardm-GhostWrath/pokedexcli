package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var cleanText []string
	cleanText = strings.Fields(strings.ToLower(text))
	return cleanText
}