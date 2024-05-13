// Nhap vao mot chuoi gom cac tu duoc phan cach bang khoang trang
// In ra ca tu do sau khi duoc loai bo cac tu trung lap va sap xep theo thu tu so va chu

package main

import (
	"fmt"
	"strings"
)

func main () {
	fmt.Printf("Nhap vao chuoi can:")

	var input string
	fmt.Scan(&input)

	// Split the input string into individual words
	words := strings.Fields(input)

	// Remove duplicate words by creating a map
	wordMap := make(map[string]bool)
	for _, word := range words {
		wordMap[word] = true
	}

	// Initialize an empty slices to store unique words
	uniqueWords := make([]string, 0, len(wordMap))

	// Extract unique words from the map
	for word := range wordMap {
		uniqueWords = append(uniqueWords, word)
	}

	// Sort the unique words alphanumerically
	fmt.Println("Unique words (sorted):")
	for _, word:= range uniqueWords {
		fmt.Println(word)
	}
}