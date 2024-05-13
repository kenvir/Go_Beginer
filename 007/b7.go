package main

import "fmt"

func main() {
	var X, Y int

	fmt.Printf("Nhap vao gia tri cua X va Y: ")

	_,err := fmt.Scan(&X, &Y)
	if err != nil {
		fmt.Printf("Invalid input. Please enter two integers!!!")
		return
	}

	// Create a 2-dimensional array
	arr := make([][]int, X)
	for i:= 0; i < X; i++ {
		arr[i] = make([]int, Y)
		for j:= 0; j < Y; j++ {
			arr[i][j] = i * j
		}
	}

	// Print the 2-dimensional array
	fmt.Println("Generated 2-dimensional array:")
	for i:= 0; i < X; i++ {
		for j:= 0; j < Y; j++ {
			fmt.Printf("%d\t", arr[i][j])
		}
		fmt.Println()
	}
}