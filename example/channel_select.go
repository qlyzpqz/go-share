package main

import (
	"fmt"
	"time"
)

func write_to_ch(ch chan int) {
	time.Sleep(2 * time.Second)
	close(ch)
}

func main() {
	var read_ch chan int = make(chan int, 1)
	var write_ch chan int = make(chan int, 1)
	var num int
	var ok bool

	go write_to_ch(read_ch)

	fmt.Printf("read_ch.len=%d, read_ch.cap=%d, write_ch.len=%d, write_ch.cap=%d\n", len(read_ch), cap(read_ch), len(write_ch), cap(write_ch))

	for i := 0; i < 2; i++ {
		select {
		// read_ch是否可读
		case num, ok = <-read_ch:
			if ok {
				fmt.Println("read_ch readabl, read_value=", num)
			} else {
				fmt.Println("read_ch closed...")
			}
		// write_ch是否可写
		case write_ch <- num:
			fmt.Println("write_ch writeable...")
		}
	}
}
