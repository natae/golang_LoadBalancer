package nataelb

import (
	"strings"
	"unicode"
)

func IsIPAddress(url string) bool {
	containsProtocol := strings.Contains(url, "http")

	containsAlphabet := false
	for _, r := range url {
		if unicode.IsLetter(r) {
			containsAlphabet = true
			break
		}
	}

	return !containsProtocol && !containsAlphabet
}
