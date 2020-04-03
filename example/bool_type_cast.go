package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b bool
	// var i int8 = 1
	// b = bool(i)
	b = true

	// fmt.Println(int(b))
	fmt.Println(*(*int)(unsafe.Pointer(&b)))
}
