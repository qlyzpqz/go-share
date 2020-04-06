package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func main() {
	var person = Person{Name: "kentpeng", Age: 29}
	var dst_person Person

	content, err := json.Marshal(person)
	if err == nil {
		fmt.Println(string(content))
	}

	err = json.Unmarshal(content, &dst_person)
	if err == nil {
		fmt.Printf("%+v\n", dst_person)
	}
}
