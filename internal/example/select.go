package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	initChan := make(chan int)
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	go func() {
		for i := 0; i < 10; i++ {
			initChan <- i
		}
	}()
	for {
		select {
		case <-initChan:
			fmt.Println(<-initChan)
		case <-ctx.Done():
			os.Exit(1)
		default:
			fmt.Println("no val")
		}
		time.Sleep(time.Second * 1)
	}
}
