package main

import (
	"fmt"
	"strings"
)

func countLines(text string) int {
	if text == "" {
		return 0
	}
	return len(strings.Split(text, "\n"))
}

func countWords(text string) int {
	if text == "" {
		return 0
	}
	return len(strings.Fields(text))
}

func countChars(text string) int {
	count := 0
	for _, r := range text {
		if r != ' ' && r != '\n' && r != '\t' {
			count++
		}
	}
	return count
}

func main() {
	text := "Hello\nWorld\nGolang"
	fmt.Printf("Nombre de lignes : %d\n", countLines(text))
	fmt.Printf("Nombre de mots : %d\n", countWords(text))
	fmt.Printf("Nombre de caractères : %d\n", countChars(text))
}
