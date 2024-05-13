// Nhap mot chuoi cac so nhi phan co 4 chu so, phan tach bang dau phay
// Kiem tra xem co chia het cho 5 khong, cac so chia se duoc in ra theo thu tu va phan tach bang dau phay

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main () {

	fmt.Printf("Nhap vao chuoi so nhi phan:")

	var input string
	fmt.Scan(&input)

	binaryNumbers := strings.Split(input, ",")
	divisibleBy5 := make([]string, 0)

	for _, binary := range binaryNumbers {
		decimal, err := strconv.ParseInt(binary, 2, 64)
		if err != nil {
			fmt.Printf("Invalid input: %s\n", binary)
			return
		}

		if decimal % 5 == 0 {
			divisibleBy5 = append(divisibleBy5, binary)
		}
	}

	fmt.Println("Binary numbers divisible by 5:")
	fmt.Println(strings.Join(divisibleBy5, ", "))
}