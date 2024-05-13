// Tim tat ca cac so trong khoang tu 100 - 300, sao co moi chu so cua so do la so chan
// Cac so thu duoc in theo trinh tu va duoc phan tach bang dau phay

package main

import "fmt"

func main() {
    // Initialize an empty slice to store the numbers
    var result []int

    // Iterate through the range of numbers from 100 to 300 (inclusive)
    for i := 100; i <= 300; i++ {
        // Check if each digit of the number is even
        if isEveryDigitEven(i) {
            // If all digits are even, append the number to the result slice
            result = append(result, i)
        }
    }

    // Print the numbers in a comma-separated sequence on a single line
    for index, num := range result {
        // Print the number
        fmt.Print(num)

        // Add a comma and space after the number if it's not the last number in the sequence
        if index < len(result)-1 {
            fmt.Print(", ")
        }
    }

    // Print a newline character at the end
    fmt.Println()
}

// Function to check if each digit of a number is even
func isEveryDigitEven(n int) bool {
    for n != 0 {
        digit := n % 10
        if digit%2 != 0 {
            return false
        }
        n /= 10
    }
    return true
}
