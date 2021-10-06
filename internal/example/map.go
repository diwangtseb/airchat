package main

import (
	"fmt"
	"strconv"
)

func _() {
	myMap := make(map[string]int)
	for i:=0;i<10;i++{
		myMap[strconv.Itoa(i)]=i
	}
	fmt.Println(myMap)
}