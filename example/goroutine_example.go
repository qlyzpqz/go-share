package main

import (
	"fmt"
	"time"
)

func goroutine(routine_id int) {
	for i := 0; i < 10; i++ {
		fmt.Println("goroutine: routine_id=", routine_id, ", idx=", i)
	}
}

func main() {
	go goroutine(1)
	go goroutine(2)

	time.Sleep(1 * time.Second)
}
