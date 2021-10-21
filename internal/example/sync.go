package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

func BorrowBooks(ctx context.Context,wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("借书完成")
			os.Exit(1)
		default:
			fmt.Println("还在借书")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx,cancel := context.WithCancel(context.Background())
	fmt.Println("借个书")
	go BorrowBooks(ctx,&wg)
	cancel()
	wg.Wait()
	time.Sleep(time.Second)
}