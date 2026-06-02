package main

func CountCharacter(str string, c rune) int {
	runes := []rune(str)

	counter := 0

	for _, char := range runes {
		if c == char {
			counter++
		}
	}
	return counter
}
