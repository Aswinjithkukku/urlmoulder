package utils

import (
	"math/rand"

	"github.com/aswinjithkukku/url-moulder/models"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func GenerateRandomString(length int, urls []models.Urls) (string, bool) {

	random := make([]byte, length)

	for i := range random {
		random[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	generatedString := string(random)

	isExist := false

	for i := range urls {
		if urls[i].SlugUrl == generatedString {
			isExist = true
			break
		}
	}

	return generatedString, isExist
}
