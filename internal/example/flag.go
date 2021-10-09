package main

import (
	"flag"
	"fmt"
)

var inputStr = flag.String("name","target","input your name")
var inputInt int

func init() {
	flag.IntVar(&inputInt,"age",0,"input your age")
}
func _() {
	flag.Parse()
	fmt.Println("name=", *inputStr)
	fmt.Println("sex=",inputInt)
}