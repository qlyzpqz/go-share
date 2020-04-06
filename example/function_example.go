package main

import (
	"fmt"
	"unsafe"
)

func say_hi(name string) {
	fmt.Println("hi, " + name)
}

func wrap_func(handle func(string)) {
	fmt.Println("before wrap_func")
	handle("xxx")
	fmt.Println("end wrap_func")
}

func main() {
	say_hi("kent")

	func_var := say_hi
	func_var("world")

	// 传入参数为函数名
	wrap_func(say_hi)

	// func变量，存储的是函数的入口地址
	// nm funcction_example | grep say_hi
	fmt.Printf("say_hi.size=%d, say_hi=%p\n", unsafe.Sizeof(say_hi), say_hi)
}
