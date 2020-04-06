package main

import (
	"fmt"
)

func main() {
	var i int
	for i = 0; i < 2; i++ {
		fmt.Println(i)
	}
	fmt.Println("---------")

	i = 0
	for i < 2 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("---------")

	i = 0
	for {
		if i >= 2 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("---------")

	var slice = []int{0, 1}
	for idx, value := range slice {
		fmt.Println("idx=", idx, ", value=", value)
	}
	fmt.Println("---------")

	var dict = map[string]string{"key1": "value1", "key2": "value2"}
	for key, value := range dict {
		fmt.Println("key=", key, ", value=", value)
	}
}

