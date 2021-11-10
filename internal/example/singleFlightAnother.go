package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"
)

func QueryKey(key string,ctx context.Context)(string,error){
	return key,nil
}

func main(){
	var sg singleflight.Group
	ctx := context.Background()
	key := "double six six six"
	ch := sg.DoChan(key,func()(interface{},error){
		rep,err := QueryKey(key,ctx)
		return rep,err
	})
	select {
		case ret := <- ch:
			fmt.Println("over",ret.Val)
	}
}