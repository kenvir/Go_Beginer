package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Nhap vao chuoi can in hoa:")

	var input []string
	fmt.Scan(&input)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, strings.ToUpper(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}


	fmt.Print("Chuoi da in hoa:")
	for _,input := range input {
		fmt.Println(input)
	}

}