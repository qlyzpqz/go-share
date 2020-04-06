package main

import (
	"fmt"
	"unsafe"
)

type MySlice struct {
	Data unsafe.Pointer
	Len int
	cap int
}

func main() {
	var array = [5]int{1, 2, 3, 4, 5}
	var slice1, slice2 []int

	fmt.Printf("before assign, slice1.addr=%p, slice2.addr=%p\n", &slice1, &slice2)

	slice1 = array[0:4]
	slice2 = slice1

	fmt.Printf("after assign, slice1.addr=%p, slice2.addr=%p\n", &slice1, &slice2)
	fmt.Printf("slice1=%+v\n", (*MySlice)(unsafe.Pointer(&slice1)))
	fmt.Printf("slice2=%+v\n", (*MySlice)(unsafe.Pointer(&slice2)))

	slice1 = append(slice1, 6)
	fmt.Printf("after append one element, slice1.addr=%p, slice2.addr=%p\n", &slice1, &slice2)
	fmt.Printf("slice1=%+v\n", (*MySlice)(unsafe.Pointer(&slice1)))
	fmt.Printf("slice2=%+v\n", (*MySlice)(unsafe.Pointer(&slice2)))

	slice1 = append(slice1, 7)
	fmt.Printf("after append two element, slice1.addr=%p, slice2.addr=%p\n", &slice1, &slice2)
	fmt.Printf("slice1=%+v\n", (*MySlice)(unsafe.Pointer(&slice1)))
	fmt.Printf("slice2=%+v\n", (*MySlice)(unsafe.Pointer(&slice2)))

}
