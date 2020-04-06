package test_pkg

import (
	"fmt"
)

func init() {
	fmt.Println("test_pkg.pkg1 init...")
}

type TestStruct struct {
	PublicField string
	privateField string
}
