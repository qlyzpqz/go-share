package main

import (
	"fmt"
)

func main() {
	var a, b = 10, 3

	fmt.Println("a=", a, ", b=", b)
	fmt.Println("a+b=", a+b, ", a-b=", a-b, ", a*b=", a*b, ", a/b=", a/b, ", a%b=", a%b)
	fmt.Println("a==10&&b==3:", (a==10&&b==3), ", a!=10||b!=3:", (a!=10||b!=3), ", !(a==10):", !(a==10))
	fmt.Println("a>=b:", (a>=b), ", a<b:", a<b)
	fmt.Printf("a=%b, b=%b, a&b=%b, a|b=%b, ^a=%b, a^b=%b, 1<<2=%b, 8>>1=%b\n", a, b, a&b, a|b, uint(^a), a^b, 1<<2, 8>>1)
}
