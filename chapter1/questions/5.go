package main

// check each condition in separate loops
func IsOneEditAway1(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}

	lengthDiff := len(s1) - len(s2)

	// if the lengths are the same, loop through the two strings and
	// check if more than one character are different
	if lengthDiff == 0 {
		differentCharacters := 0

		for index, s1Character := range s1 {
			s2Character := []rune(s2)[index]
			if s1Character != s2Character {
				differentCharacters++
			}
		}

		return differentCharacters <= 1
	}

	// if the lengths are different, loop through the two strings.
	// when a different character is encountered, skip the character of the
	// longer string and continue looping.
	// check if more than one character needs skipping
	var longer []rune
	var shorter []rune
	if lengthDiff == 1 {
		longer = []rune(s1)
		shorter = []rune(s2)
	} else if lengthDiff == -1 {
		longer = []rune(s2)
		shorter = []rune(s1)
	} else {
		// more than one edit away if length differs by 2 or more
		return false
	}

	longerPointer := 0
	differentCharacters := 0
	for _, shorterCharacter := range shorter {
		longerCharacter := longer[longerPointer]

		if longerCharacter == shorterCharacter {
			longerPointer++
		} else {
			longerPointer += 2
			differentCharacters++
		}

		if differentCharacters > 1 {
			return false
		}
	}

	return true
}


// check each condition in one loop
func IsOneEditAway2(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}

	lengthDiff := len(s1) - len(s2)

	if lengthDiff > 1 || lengthDiff < -1 {
		return false
	}

	s1Pointer := 0
	s2Pointer := 0

	diffFound := false

	for {
		if s1Pointer == len(s1) || s2Pointer == len(s2) {
			break
		}

		s1Character := []rune(s1)[s1Pointer]
		s2Character := []rune(s2)[s2Pointer]

		if s1Character == s2Character {
			s1Pointer++
			s2Pointer++
			continue
		}

		// when the length is the same, simply mark the difference and continue
		if lengthDiff == 0 {
			if diffFound {
				// this is the second difference found
				return false
			}

			diffFound = true
			s1Pointer++
			s2Pointer++
			continue
		}

		// when the length is not the same, keep the shorter string pointer at the same position,
		// and move the longer string pointer +1.
		if lengthDiff != 0 {
			if diffFound {
				// this is the second difference found
				return false
			}

			diffFound = true

			if lengthDiff > 0 {
				// s1 is longer than s2
				s1Pointer++
			} else {
				// s2 is longer than s1
				s2Pointer++
			}
		}
	}

	return true
}
