package main

import (
	"fmt"
)

func test_func() {
	fmt.Println("test_func begin")
	panic("unexpected error...")
	fmt.Println("test_func end")
}

func main() {
	defer func() {
		fmt.Println("defer one")
	}()
	defer func() {
		fmt.Println("defer two")
		errmsg := recover()
		if errmsg != nil {
			fmt.Println("catch panic,", errmsg)
		}
	}()
	defer func() {
		fmt.Println("defer three")
		panic("panic at defer")
	}()
	defer func() {
		errmsg := recover()
		if errmsg != nil {
			fmt.Println("catch panic,", errmsg)
		}
	}()

	test_func()
}
