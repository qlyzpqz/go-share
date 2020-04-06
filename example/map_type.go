package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var dict map[string]int

	fmt.Printf("map.size=%d, map=%p\n", unsafe.Sizeof(dict), dict)
	dict = make(map[string]int, 0)
	fmt.Printf("map.size=%d, map=%p\n", unsafe.Sizeof(dict), dict)
}
