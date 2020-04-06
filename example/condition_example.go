package main

import (
	"fmt"
)

func main() {
	var a int = 5
	var b interface{}

	// if..else if..else
	if a <= 0 {
		fmt.Println("a <= 0")
	} else if a > 0 && a <= 5 {
		fmt.Println("a >= 0 && a <= 5")
	} else {
		fmt.Println("a > 5")
	}

	b = a

	// switch不带变量
	switch {
	case a <= 0:
		fmt.Println("a <= 0")
	case a > 0 && a <= 5:
		fmt.Println("a >= 0 && a <= 5")
	case a > 5:
		fmt.Println("a > 5")
	default:
		fmt.Println("default...")
	}

	// switch带变量
	switch a {
	case 1:
		fmt.Println("a==1")
	case 5:
		fmt.Println("a==5")
	default:
		fmt.Println("default...")
	}

	// 对类型进行switch
	switch b.(type) {
	case int:
		fmt.Println("b.(type) == int")
	case string:
		fmt.Println("b.(type) == string")
	default:
		fmt.Println("default...")
	}
}
