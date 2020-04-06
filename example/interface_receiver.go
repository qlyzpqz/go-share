package main

import (
	"fmt"
)

type TestInterface interface {
	fun1()
	fun2()
}

type TestStruct struct {
}

// 定义一个值receiver的方法
func (this TestStruct) fun1() {
	fmt.Println("fun1, value receiver")
}

// 定义一个指针receiver的方法
func (this *TestStruct) fun2() {
	fmt.Println("fun1, pointer receiver")
}

func main() {
	var st1 TestInterface = &TestStruct{}

	st1.fun1()
	st1.fun2()

	// 编译会报告
	// 对象在转为interface时，会做数据复制，指针接收者方法是可以修改对象的值的，如果允许TestStruct{}赋值为TestInterface，就会出现指针接收者方法修改的值实际是数据副本，没有真正对原数据产生副作用,语义上不符合设计
	var st2 TestInterface = TestStruct{}
}
