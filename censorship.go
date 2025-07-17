package main

import (
	"strings"
)

func badwordsMap() map[string]string {
	return map[string]string{
	"kerfuffle": "****",
	"sharbert":"****",
	"fornax": "****",
	}
}

func censorBadWords(msg string) string {
	badwords := badwordsMap()
	lowercasewords := strings.Split(strings.ToLower(msg), " ")
	normalcasewords := strings.Split(msg, " ")
	for i, word := range lowercasewords {
		rep, exists := badwords[word]
		if exists {
			normalcasewords[i] = rep
		}
	}
	return strings.Join(normalcasewords, " ")
}
