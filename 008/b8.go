// Nhap tu, duoc phan tach bang dau phay
// In ra cac tu khi chung da duoc sap xep a-z

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
    // Prompt the user to enter a comma-separated sequence of words
    fmt.Printf("Enter a comma-separated sequence of words:")

    // Read the input from the console
	var input string
	fmt.Scan(&input)

    // Split the input string into individual words
    words := strings.Split(input, ",")

    // Trim any leading or trailing whitespace from each word
	for i, word := range words {
		words[i] = strings.TrimSpace(word)
	}

    // Sort the words alphabetically
    sort.Strings(words)

    // Print the sorted words in a comma-separated sequence
    fmt.Println("Sorted words:", strings.Join(words, ", "))
}
