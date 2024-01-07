package main

import "fmt"

// replace all spaces with %20
// example
// s: "Mr John Smith      "
// trueLength: 13 (length up to the last non space character)
func URLify(s string, trueLength int) string {
	characters := []rune(s)
	outputPointer := len(s) - 1

	for inputPointer := trueLength - 1; inputPointer >= 0; inputPointer-- {
		fmt.Printf("output: %d, input: %d", outputPointer, inputPointer)
		fmt.Printf("~%s~\n", string(characters))
		if characters[inputPointer] == ' ' {
			characters[outputPointer] = '0'
			characters[outputPointer-1] = '2'
			characters[outputPointer-2] = '%'
			outputPointer -= 3
		} else {
			characters[outputPointer] = characters[inputPointer]
			outputPointer--
		}
	}

	if outputPointer > 0 {
		characters = characters[outputPointer+1:]
	}

	return string(characters)
}
