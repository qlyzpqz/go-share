package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b bool

	b = true

	fmt.Println(b)
	fmt.Println(unsafe.Sizeof(false))
}
