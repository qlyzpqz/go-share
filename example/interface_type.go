package main

import (
	"fmt"
	"unsafe"
)

// 定义一个Animal接口
type Animal interface {
	Speak () string
}

type Cat struct {
	Name string
}

// 实现了Animal接口
func (this Cat) Speak() string {
	return "Hi, my name is " + this.Name + ", I'm a cat"
}

type Mouse struct {
	Name string
}

// 实现了Animal接口
func (this Mouse) Speak() string {
	return "Hi, my name is " + this.Name + ", I'm a mouse"
}

// 参数为Animal接口类型
func speak(animal Animal) {
	fmt.Println(animal.Speak())
}

func main() {
	var cat = Cat{Name: "Tom"}
	var mouse = Mouse{Name: "Jerry"}
	var animal Animal

	// 参数可以传入Cat类型,因为Cat实现了Speak方法
	speak(cat)

	// 参数可以传入Mouse类型,因为Mouse实现了Speak方法
	speak(mouse)

	fmt.Printf("animal.size=%d\n", unsafe.Sizeof(animal))
}

