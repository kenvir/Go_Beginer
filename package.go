package main

import (
	"fmt"
)

// Khai bao bien va tao ham add tinh tong

func add(x, y int) int{
	return x + y
}

// Ham loop, mang gom cac phan tu string, vong lap va in ra tung phan tu trong mang
func loop() {
	program := []string{"JS", "Java", "Go"}

	for i := 0; i < len(program); i++ {
		fmt.Println(program[i])
	}
}



func main() {
	fmt.Println(add(2, 9))
	loop()
}