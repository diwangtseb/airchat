package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	HaveBuffer()
	NoBuffer()
	SingleChan(make(chan<- string, 1), make(<-chan string, 1))
	chanEle := make(chan string)
	for i := 0; i < 10; i++ {
		go func(i int) {
			chanEle <- strconv.Itoa(i)
		}(i)
	}
	ChanSelect(chanEle)
}

func HaveBuffer() {
	message := make(chan int, 1)
	message <- 1
	msg := <-message
	fmt.Println(msg)
}

func NoBuffer() {
	message := make(chan string)
	go func() {
		message <- "1"
	}()
	msg := <-message
	fmt.Println(msg)
}

func SingleChan(ping chan<- string, pong <-chan string) {
	ping <- "a"
}
func ChanSelect(ping chan string) {
	for {
		select {
		case msg := <-ping:
			fmt.Println(msg)
		default:
			fmt.Println("no arg")
			time.Sleep(time.Second)
		}

	}
}
