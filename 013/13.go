// Nhap vao mot chuoi, tinh so ki ty chu va so

package main

import (
	"fmt"
	"unicode"
)

func main() {
    // Prompt the user to enter a sentence
    fmt.Println("Enter a sentence:")

    // Read the input from the console
    var sentence string
    fmt.Scanln(&sentence)

    // Initialize counters for letters and digits
    numLetters := 0
    numDigits := 0

    // Iterate through each character in the sentence
    for _, char := range sentence {
        // Check if the character is a letter
        if unicode.IsLetter(char) {
            numLetters++
        }
        // Check if the character is a digit
        if unicode.IsDigit(char) {
            numDigits++
        }
    }

    // Print the number of letters and digits
    fmt.Printf("Number of letters: %d\n", numLetters)
    fmt.Printf("Number of digits: %d\n", numDigits)
}
