package main

import (
	"fmt"
)

func is_negative(num int) bool {
	if num < 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(is_negative(-1))
}
