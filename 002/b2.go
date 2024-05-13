// Tinh giai thua cua mot so cho truoc.
// Cac ket qua phai duoc in theo trinh tu, duoc phan tach bang dau phay tren mot dong.

package main

import (
	"fmt"
	"log"
)

func Factorial(input int) (uint64, error) {
	var factorial uint64 = 1

	// Check enter input
	if input < 0 {
		return 0, fmt.Errorf("factorial for negativ values is not allowed")
	} else if input == 0 {
		return 1, nil
	}

  	// All if true => factorial(i)
	for i := 1; i <= input; i++ {
		factorial *= uint64(i)
	}

	return factorial, nil
}

func main() {
	var input int

	fmt.Printf("Nhap vap so can tinh giai thua: ")
	_,err := fmt.Scanln(&input)

	if err != nil {
		log.Fatalf("Please enter number")
	}

	result, err := Factorial(input)

	if err != nil {
		log.Fatalf("Error for input %v: %v", input, err)
	}

	fmt.Printf("Factorial of %d = %d", input, result)
}
