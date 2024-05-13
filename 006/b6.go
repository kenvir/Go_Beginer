// Tinh Q voi C=50 and H=50
// Su dung cong thuc: Q = Square root of [(2 * C * D)/H]

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func calculateQ(D float64) float64 {
    // Fixed values of C and H
    C := 50.0
    H := 30.0

    // Calculate Q using the given formula
    return math.Sqrt((2 * C * D) / H)
}

func main() {
    // Prompt the user to enter the values of D
    fmt.Printf("Enter the values of D (comma-separated): ")

    // Read the input from the console
    var input string
    fmt.Scanln(&input)

    // Split the input string into individual values of D
    values := strings.Split(input, ",")

    // Initialize an empty slice to store the results
    var results []string

    // Convert each value of D to float64, calculate Q, and store the result
    for _, val := range values {
        d, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
        if err != nil {
            fmt.Printf("Invalid input: %s\n", val)
            return
        }
        q := calculateQ(d)
        results = append(results, fmt.Sprintf("%.2f", q))
    }

    // Print the results
    fmt.Println("Q values:", strings.Join(results, ", "))
}
