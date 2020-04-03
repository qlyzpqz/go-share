package main

import (
	"fmt"
)

func main() {
	var i8 int8 = 8
	var i32 int32 = 32
	
	// 失败的转换
	// i32 = i8

	// 成功的转换
	i32 = int32(i8)

	fmt.Println(i32)
}
