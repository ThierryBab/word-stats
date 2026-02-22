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
}
