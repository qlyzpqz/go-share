package main

import (
	"fmt"
)

func main() {
	var i int8
	var j int8 = 86

	i = j + j

	// 溢出变为负值
	fmt.Printf("i=%d, j=%d\n", i, j)
	fmt.Printf("i=%#x, j=%#x\n", uint8(i), uint8(j))
}
