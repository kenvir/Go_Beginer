// Tim tat ca cac so chia het cho 7 nhung khong phai boi so cua 5 trong khoang tu 2000 - 3200.
// Cac so thu duoc phai in theo thu tu, phan tach bang dau phay tren mot dong

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FindNum (low, hight int) string{
	// Tao mot mang numbers co kieu string
	var numbers []string

	// Lap qua cac phan tu tu low - hight
	for i := low; i <= hight; i++ {
		// Kiem tra dieu kien i chia het cho 7 va khong chia het cho 5
		if i % 7 == 0 && i % 5 != 0 {
			// Chuyen doi i thanh kieu chuoi va them i vao slice numbers
			numbers = append(numbers, strconv.Itoa(i))
		}
	}
	// Tra ve chuoi duoc them vao numbers va phan cach nhau bang dau phay
	return strings.Join(numbers, ",")
}

func main () {
	// Khoi tao 2 gia tri cho 2 bien low va hight
	res := FindNum(2000, 3200)
	// In ra ket qua cua ham FindNum khi da them 2 bien
	fmt.Println(res)
}