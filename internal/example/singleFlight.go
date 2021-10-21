package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var count int64

func GetArticle(id int64) (string,error){
	atomic.AddInt64(&count,1)
	time.Sleep(time.Millisecond*time.Duration(count))
	return fmt.Sprintf("%d's article is xxx",id),nil
}

func SingleFligtGetArticle(sg *singleflight.Group,id int64)(string ,error) {
	v,err,_:= sg.Do(strconv.FormatInt(id,10),func()(interface{}, error){
		return GetArticle(id)
	})
	if err != nil {
		return "",err
	}
	return v.(string),nil
}

func _(){
	time.AfterFunc(1*time.Second, func() {
		atomic.AddInt64(&count,-count)
	})
	var (
		wg sync.WaitGroup
		now = time.Now()
		n = 1000
		sg = &singleflight.Group{}
	)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			//res,_ := GetArticle(1)
			res,_ := SingleFligtGetArticle(sg,1)
			if res != "1's article is xxx"{
				panic("err")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("同时发起%d次请求，耗时:%s\n",n,time.Since(now))
}
