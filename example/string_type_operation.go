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
	var str1, str2 string

	str1 = "str1"
	str2 = str1

	fmt.Printf("str1 struct, %+v\n", (*MyString)(unsafe.Pointer(&str1)))
	fmt.Printf("str2 struct, %+v\n", (*MyString)(unsafe.Pointer(&str2)))

	str2 = str2 + ".."
	fmt.Printf("str1 struct, after operation, %+v\n", (*MyString)(unsafe.Pointer(&str1)))
	fmt.Printf("str2 struct, after operation, %+v\n", (*MyString)(unsafe.Pointer(&str2)))
}
