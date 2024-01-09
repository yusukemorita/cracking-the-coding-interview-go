package main

import (
	"strconv"
	"strings"
)

// palindrome: word or phrase that is the same
// read backwards and forwards
// permutation: same characters, different order
func IsPalindromePermutation1(s string) bool {
	characters := map[rune]uint{}

	// count the number of all characters in the string
	for _, character := range s {
		// ignore spaces
		if character != ' ' {
			characters[character]++
		}
	}

	// only one character with and odd count is allowed
	// all other characters must have an even count
	oddCharacterCount := 0
	for _, count := range characters {
		if count%2 != 0 {
			oddCharacterCount++
		}
		if oddCharacterCount > 1 {
			return false
		}
	}

	return true
}

// use a bit vector to decrease space usage
// assume only lowercase alphabet + space characters (27)
func IsPalindromePermutation2(s string) bool {
	// only 26 bytes are needed, so int32 is enough
	// (spaces are ignored, so only need to store counts of the 26 letters of the alphabet)
	// however, strconv.FormatInt only accepts a int64, so int64 is used instead of int32
	var characters int64 = 0

	// count the number of all characters in the string
	for _, character := range s {
		// ignore spaces
		if character != ' ' {
			// assign character a number, starting 'a' with 0
			val := character - 'a'
			// use XOR bitwise operator, so that
			// - if the count for this character is 1, then the total count so far (not including this one) is odd.
			//   since including this character it will be even again, we want the bit for this character to become 0.
			// - all other bits will remain unchanged, as 0 ^ 0 = 0, and 1 ^ 0 = 1
			characters = characters ^ (1 << int(val))
		}
	}

	// when converted to binary, the number of 1s must be 0~1,
	// and everything else must be 0.
	binaryString := strconv.FormatInt(characters, 2)
	numberOfOnes := len(strings.Replace(binaryString, "0", "", -1))

	return numberOfOnes <= 1
}
