package main

import (
	"fmt"
	"runtime"
)

func _() {
	go func() {
		fmt.Println("i am go another go routine")
	}()
	go firstFunc()
	defer fmt.Println(fact(7))
	fmt.Println(runtime.NumGoroutine())
	fmt.Println("i am main go routine")
}
func firstFunc() {
	go func() {
		fmt.Println("i am firstFunc go routine")
		fact(100)
		intSeq(100)
	}()
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func intSeq(n int) func() int {
	if n > 0 {
		return func() int {
			return 1
		}
	}
	return func() int {
		return 0
	}
}
