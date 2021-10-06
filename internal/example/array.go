package main

import "fmt"

func _() {
	myArrary := [3]int{1,2,3}
	for index,value := range myArrary{
		fmt.Println(index,value)
	}
}
