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

func main() {
	text := "Hello\nWorld\nGolang"
	fmt.Printf("Nombre de lignes : %d\n", countLines(text))
}
