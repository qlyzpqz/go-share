package main

import (
	"time"
)

func goroutine_sleep() {
	time.Sleep(60 * time.Second);
}

func main() {
	// 创建100w个goroutine
	for i := 0; i < 1000000; i++ {
		go goroutine_sleep()
	}

	time.Sleep(61 * time.Second);
}
