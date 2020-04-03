package main

import (
	"fmt"
	"unsafe"
)

type MyString struct {
	Data unsafe.Pointer
	Len int
}

func main() {
	var str string = "Hello World............."

	fmt.Printf("string.size=%d, len(string)=%d\n", unsafe.Sizeof(str), len(str))

	ms := *(*MyString)(unsafe.Pointer(&str))
	fmt.Printf("%s: %d\n", string(*(*[]byte)(ms.Data)), ms.Len)
}
