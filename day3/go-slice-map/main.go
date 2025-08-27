package main

import (
	"fmt"
	"strings"
)


func countWordFrequency(text string) map[string]int {
	lowerText := strings.ToLower(text)

	words := strings.Fields(lowerText)

	wordCounts := make(map[string]int)

	for _, word := range words {
		wordCounts[word]++
	}

	return wordCounts
}

func main() {
	sentence := "go is awesome go is fast and go is fun"
	frequencyMap := countWordFrequency(sentence)

	fmt.Println("Tần suất xuất hiện của các từ:")
	for word, count := range frequencyMap {
		fmt.Printf("- %s: %d\n", word, count)
	}
}