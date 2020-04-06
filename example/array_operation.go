package main

import (
	"fmt"
)

func main() {
	var arr1 [16]byte
	var arr2 [16]byte

	fmt.Printf("before assign, arr1=%p, arr2=%p\n", &arr1, &arr2)
	arr1 = arr2
	fmt.Printf("after assign,  arr1=%p, arr2=%p\n", &arr1, &arr2)
	// fmt.Println(len(arr1))
}
