package main

import (
	"strconv"
	"strings"
)

func CompressString(s string) string {
	if len(s) <= 2 {
		return s
	}

	var compressed strings.Builder
	characters := []rune(s)

	var lastCharacterCount uint64 = 1
	var lastCharacter rune

	for index, character := range characters {
		if index > 0 && character == characters[index-1] {
			lastCharacterCount++
		} else {
			if index > 0 {
				compressed.Write([]byte(string(lastCharacter) + strconv.FormatUint(lastCharacterCount, 10)))
			}

			lastCharacterCount = 1
			lastCharacter = character
		}
	}

	compressed.Write([]byte(string(lastCharacter) + strconv.FormatUint(lastCharacterCount, 10)))

	if compressed.Len() < len(s) {
		return compressed.String()
	} else {
		return s
	}
}
