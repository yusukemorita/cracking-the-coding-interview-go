package main

import "strings"

// e.g. "erbottlewat" is a rotation of "waterbottle"
func IsRotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if s1 == s2 {
		return true
	}

	return IsSubString(s1 + s1, s2)
}

// only allowed to be called once
func IsSubString(big, small string) bool {
	return strings.Contains(big, small)
}