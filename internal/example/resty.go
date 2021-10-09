package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)
type RequestModel struct {
	Id int8
	Name string
}

func main() {
	reqm := RequestModel{Id: 1,Name: "66"}
	client := resty.New()
	resp, _ := client.R().SetBody(reqm).Post("http://localhost:8080/ping/post")
	fmt.Println(resp)
}