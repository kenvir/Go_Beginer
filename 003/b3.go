// Voi mot so nguyen n cho truoc, tao mot map chua (i, i*i) sao cho so nguyen nam trong khoang tu 1 - n.
// Chuong trinh se in map voi gia tri cua dai dien

package main

import (
	"fmt"
	"log"
)

func NumSquared(n int) map[int]int {
	var numbers = make(map[int]int, n)

	for i := 1; i <= n; i++ {
		numbers[i] = i * i
	}

	return numbers
} 

func main() {
	var n int
	fmt.Printf("Nhap vao n: ")

	_,err := fmt.Scanln(&n)

	if err != nil {
		log.Fatal("Error occured: ", err)
	}

	fmt.Printf("%v", NumSquared(n))
}