package main

import (
	"fmt"
)

func main() {
	var dict map[string]string
	var dict_copy map[string]string

	dict = make(map[string]string)

	// 增加元素
	dict["key"] = "value"
	fmt.Printf("after insert, dict=%+v\n", dict)

	// 判断元素是否存在
	if _, exist := dict["key"]; exist {
		fmt.Println("key exist in dict")
	}

	// 删除元素
	delete(dict, "key")
	fmt.Printf("after delete, dict=%+v\n", dict)

	dict_copy = dict
	fmt.Printf("&dict=%p, dict=%p\n", &dict, dict)
	fmt.Printf("&dict_copy=%p, dict_copy=%p\n", &dict_copy, dict_copy)
}
