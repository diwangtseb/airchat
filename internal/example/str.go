package main

import "fmt"

func _(){
	var a  = "北京市"
	fmt.Println(string([]rune(a)[len([]rune(a))-1:]))
}
