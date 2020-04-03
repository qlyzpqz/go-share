package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b byte = 'a'
	var r rune = '世'

	// "世"的unicode为\u4e16
	fmt.Printf("b=%#x, r=%#x\n", b, r)

	fmt.Printf("b.size=%d, r.size=%d\n", unsafe.Sizeof(b), unsafe.Sizeof(r))
}
