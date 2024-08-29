package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "My name is Joe and My Father's name is also Joe"
	result := findWordsWithHighestFrequency(input)
	fmt.Println(result)
}

func findWordsWithHighestFrequency(input string) []string {
	words := strings.Fields(input)

	frequencyMap := make(map[string]int)
	var orderSlice []string

	for _, word := range words {
		repetation := frequencyMap[word]
		frequencyMap[word] = repetation + 1
		if val, exists := frequencyMap[word]; exists && val < 2 {
			orderSlice = append(orderSlice, word)
		}
	}

	maxFrequency := 0
	for _, freq := range frequencyMap {
		if freq > maxFrequency {
			maxFrequency = freq
		}
	}

	// Collect words with the highest frequency, preserving order
	var result []string
	for _, word := range orderSlice {
		if frequencyMap[word] == maxFrequency {
			result = append(result, word)
		}
	}

	return result
}
