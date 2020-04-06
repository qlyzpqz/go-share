package main

import (
	"fmt"
	// "time"
)

func read_from_ch(ch <-chan int) {
	fmt.Println("waiting for channel readable..")
	// 会阻塞，等待可读
	val := <-ch
	fmt.Println("read_from_ch: ", val)
}

func main () {
	var nobuf_ch chan int = make(chan int)
	var buf_ch chan int = make(chan int, 2)

	// 如果没有reader，就不能往channel里写入数据,因为阻塞会导致死锁
	go read_from_ch(nobuf_ch)
	// time.Sleep(1 * time.Second)
	nobuf_ch <- 1

	// 有缓冲，写入没人读取，也不阻塞
	buf_ch <- 1
	buf_ch <- 2
	// 写满会阻塞, 没人读取会导致死锁
	// buf_ch <- 3

	// 只读channle，不能写入数据
	// var read_ch <-chan int = make(<-chan int, 10)
	// read_ch <- 10

	// 只写channel，不能读取数据
	// var write_ch chan<- int = make(chan<- int, 10)
	// var tmp int
	// tmp = <-write_ch
}
