package stringutil

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadString reads a string from console input
func ReadString() string {
    fmt.Println("Enter a string:")
    reader := bufio.NewReader((os.Stdin))
    input, _ := reader.ReadString('\n')
    return strings.TrimSpace(input)
}

// PrintString returns the string in upper case
func PrintString(str string) string {
    return strings.ToUpper(str)
}
