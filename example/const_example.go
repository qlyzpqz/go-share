package main

import (
	"fmt"
)

func main() {
	const a = 5
	const (
		b = 6
		c = 7
	)

	fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
}
