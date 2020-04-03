package main

import (
	"fmt"
)

func main() {
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64
	var i int
	var ui int

	i8 = 127
	ui8 = 255

	i16 = 16
	ui16 = 16
	i32 = 32
	ui32 = 32
	i64 = 64
	ui64 = 64
	i = 1234
	ui = 1234

	fmt.Printf("i8=%d, i16=%d, i32=%d, i64=%d, i=%d\n", i8, i16, i32, i64, i)
	fmt.Printf("ui8=%d, ui16=%d, ui32=%d, ui64=%d, i=%d\n", ui8, ui16, ui32, ui64, ui)
}
