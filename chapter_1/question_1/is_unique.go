package main

// implementation using a hash
func AreRunesUnique1(s string) bool {
	encounteredRunes := make(map[rune]bool)

	for _, r := range s {
		ok, _ := encounteredRunes[r]
		if ok {
			return false
		}

		encounteredRunes[r] = true
	}

	return true
}

// assume ASCII (128 characters, 7 bytes) characters only
// instead of using a map, use a slice of length 128
func AreRunesUnique2(s string) bool {
	checker := make([]bool, 128)

	for _, r := range s {
		asciiCode := int(r)

		value := checker[asciiCode]
		if value {
			return false
		}

		checker[asciiCode] = true
	}

	return true
}

// assume lowercase alphabet ASCII characters only (26)
// instead of using a slice, use a bit vector
func AreRunesUnique3(s string) bool {
	var checker int32 = 0
	for _, character := range s {
		val := character - 'a'
		if (checker & (1 << val)) > 0 {
			// If bit already set, return false
			return false
		}
		// Set the bit for this character
		checker |= (1 << val)
	}
	return true
}
