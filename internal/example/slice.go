 package main

 import "fmt"

 func _() {
	 mySlice := []int{}
	 for i:=0;i<10;i++{
		 mySlice = append(mySlice, i)
	 }
	fmt.Println(mySlice)
 }