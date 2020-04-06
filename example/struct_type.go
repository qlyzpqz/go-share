package main

import (
	"fmt"
	"unsafe"
)

type MyStruct struct {
	Field1 int
	Field2 string
}

func (this MyStruct) Print() {
	fmt.Printf("field1=%d, field2=%s\n", this.Field1, this.Field2)
}

// 无法对原结构体产生副作用
func (this MyStruct) SetField1(value int) {
	// 打印方法内部结构体的内存地址
	fmt.Printf("value receiver, inner MyStruct address : %p\n", &this)
	this.Field1 = value
}

// 可以对原结构体产生副作用
func (this *MyStruct) SetField1ByPointer(value int) {
	// 打印方法内部结构体的内存地址
	fmt.Printf("pointer receiver, inner MyStruct address : %p\n", this)
	this.Field1 = value
}

func main() {
	var st1 = MyStruct{Field1: 10, Field2: "hello world"}

	// 打印结构体占用内存大小
	fmt.Println(unsafe.Sizeof(st1))

	// 结构体在函数外的内存地址
	fmt.Printf("outer MyStruct address : %p\n", &st1)
	st1.Print()
	st1.SetField1(11)
	st1.Print()
	st1.SetField1ByPointer(12)
	st1.Print()

	var pst2 = &MyStruct{Field1: 10, Field2: "hello world"}
	fmt.Printf("outer MyStruct address : %p\n", pst2)
	pst2.Print()
	pst2.SetField1(11)
	pst2.Print()
	pst2.SetField1ByPointer(12)
	pst2.Print()
}
