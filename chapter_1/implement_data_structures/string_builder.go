package main

type StringBuilder struct {
	characterList ArrayList[rune]
}

func NewStringBuilder() StringBuilder {
	return StringBuilder{
		characterList:  ArrayList[rune]{},
	}
}

func (builder *StringBuilder) addString(s string) {
	for _, runeInString := range s {
		builder.characterList.push(runeInString)
	}
}

func (builder StringBuilder) toString() string {
	return string(builder.characterList.array)
}
