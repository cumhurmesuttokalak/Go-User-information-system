package utils

import "fmt"

func IsEmpty(data string) bool {
	fmt.Println("True")
	return len(data) == 0
}
func IsNotEmpty(data string) bool {
	fmt.Println("False")
	return len(data) != 0
}
