package main

import (
	"slices"
)

// permutation: same characters, different order
// loops through one string twice, one string once so O(3N) = O(N)
func IsPermutation1(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	characterCount1 := map[rune]uint{}
	for _, character := range s1 {
		characterCount1[character] += 1
	}

	characterCount2 := map[rune]uint{}
	for _, character := range s2 {
		characterCount2[character] += 1
	}

	for character := range characterCount1 {
		if characterCount1[character] != characterCount2[character] {
			return false
		}
	}

	return true
}

// sort the characters, then compare the strings
// sorting is O(NlogN)
func IsPermutation2(s1 string, s2 string) bool {
	return sortCharacters(s1) == sortCharacters(s2)
}

func sortCharacters(s string) string {
	characters := []rune(s)
	slices.Sort(characters)
	return string(characters)
}