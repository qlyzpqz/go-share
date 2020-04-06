package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var arr [10]byte = [10]byte{1, 2}

	arr[2] = 3

	fmt.Printf("array.size=%d, array.len=%d\n", unsafe.Sizeof(arr), len(arr))
	fmt.Printf("array=%+v\n", arr)
}
