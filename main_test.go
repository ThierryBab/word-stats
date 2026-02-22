package main

import "testing"

func TestCountLines(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"Hello\nWorld", 2},
		{"One line", 1},
		{"", 0},
	}

	for _, tt := range tests {
		if got := countLines(tt.input); got != tt.want {
			t.Errorf("countLines(%q) = %d; want %d", tt.input, got, tt.want)
		}
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"Hello World Golang", 3},
		{"Single", 1},
		{"Hello   World", 2},
		{"", 0},
	}

	for _, tt := range tests {
		if got := countWords(tt.input); got != tt.want {
			t.Errorf("countWords(%q) = %d; want %d", tt.input, got, tt.want)
		}
	}
}

func TestCountChars(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"Hi Go", 4},
		{"Hello", 5},
		{"A B\nC", 3},
		{"A\tB\nC", 3},
		{"", 0},
	}

	for _, tt := range tests {
		if got := countChars(tt.input); got != tt.want {
			t.Errorf("countChars(%q) = %d; want %d", tt.input, got, tt.want)
		}
	}
}

func TestMainExecution(t *testing.T) {
	main()
}
