package main

import (
	"fmt"
	"github.com/qlyzpqz/go-share/example/test_pkg"
	_ "github.com/qlyzpqz/go-share/example/test_pkg/sub_pkg"
)

func init() {
	fmt.Println("main package init...")
}

func main() {
	var ts test_pkg.TestStruct
	ts.PublicField = "aaa"
	// ts.privateField = "bbb"

	fmt.Printf("%+v\n", ts)
}
