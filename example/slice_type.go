package main

import (
	"fmt"
	"unsafe"
)

type MySlice struct {
	Data unsafe.Pointer
	Len int
	Cap int
}

func main() {
	var slice []byte
	var arr = [6]byte{1, 2, 3, 4, 5, 6}
	slice = arr[1:3]

	fmt.Printf("slice.size=%d\n", unsafe.Sizeof(slice))
	fmt.Printf("slice=%+v\n", *(*MySlice)(unsafe.Pointer(&slice)))
	fmt.Printf("array.addr=%p\n", &arr)
}
