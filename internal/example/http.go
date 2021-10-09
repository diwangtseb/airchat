package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func _() {
	url := "http://localhost:8080/ping/post"
	bodys := []byte("{\"id\":1,\"name\":\"TestOK\"}")
	req, _:= http.NewRequest("POST", url, bytes.NewBuffer(bodys))
	client := &http.Client{}
	resp,_ := client.Do(req)
	defer resp.Body.Close()
	body,_:= ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}