// Chap nhan mot chuoi cac so duoc phan tach voi nhau bang dau phay tu console va tao mot silce tu chung.
// Tra lai slice.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Slice(input string) []int{
	numbers := strings.Split(input, ",")

	var parseNum []int

	for _, numStr := range numbers {
		num, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			fmt.Printf("Invalid input: %s\n", numStr)
		} 
		
		parseNum = append(parseNum, num)
	}
	return parseNum
}

func main() {
	fmt.Printf("Nhap vao chuoi cac so: ")

	var n string
	fmt.Scanln(&n)

	fmt.Printf("%v", Slice(n))
}