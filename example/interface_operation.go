package main

import (
	"fmt"
	"unsafe"
)

type MyIface struct {
	itab unsafe.Pointer
	data unsafe.Pointer
}

type Interface1 interface {
	fun1()
}

type Interface2 interface {
	fun2()
}

type TestStruct struct {
	Name string
}

func (this TestStruct) fun1() {
	fmt.Println("fun1")
}

func (this *TestStruct) fun2() {
	fmt.Println("fun2")
}

func main() {
	var st1 = TestStruct{}
	var if1 Interface1
	var if2 Interface1

	// 对象转为interface
	fmt.Printf("st1.address=%p\n", &st1)
	if1 = st1
	fmt.Printf("iface=%+v\n", *((*MyIface)(unsafe.Pointer(&if1))))

	fmt.Printf("st1.address=%p\n", &st1)
	if2 = &st1
	fmt.Printf("iface=%+v\n", *((*MyIface)(unsafe.Pointer(&if2))))

	// 类型断言
	if _, ok := if1.(TestStruct); ok {
		fmt.Println("if1 is type of TestStruct")
	}
}
