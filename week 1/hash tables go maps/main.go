package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	freq := make(map[string]int)

	words := strings.Fields(text)

	for _, word := range words {
		freq[word]++
	}

	return freq
}

func main() {
	text := "Go is fast and Go is powerful"

	result := wordFrequency(text)

	for word, count := range result {
		fmt.Println(word, ":", count)
	}
}