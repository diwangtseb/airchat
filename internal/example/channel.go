package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func _() {
	// HaveBuffer()
	// NoBuffer()
	// SingleChan(make(chan<- string, 1), make(<-chan string, 1))
	// chanEle := make(chan string)
	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		chanEle <- strconv.Itoa(i)
	// 	}(i)
	// }
	//ChanSelect(chanEle)
	//TimeOutChannel()
	//CloseChannel()
	//RangeChannel()
	//TickerChannel()
	//WorkerChannel()
	Atomic()
}

func HaveBuffer() {
	message := make(chan int, 1)
	message <- 1
	msg := <-message
	fmt.Println(msg)
}

func NoBuffer() {
	message := make(chan string)
	go func() {
		message <- "1"
	}()
	msg := <-message
	fmt.Println(msg)
}

func SingleChan(ping chan<- string, pong <-chan string) {
	ping <- "a"
}
func ChanSelect(ping chan string) {
	for {
		select {
		case msg := <-ping:
			fmt.Println(msg)
		default:
			fmt.Println("no arg")
			time.Sleep(time.Second)
		}

	}
}
func TimeOutChannel() {
	c1 := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- 1
	}()
	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout second *1")
	}
}

func CloseChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			fmt.Println(time.Now())
			msg, more := <-jobs
			if more {
				fmt.Println(msg, "receive")
			} else {
				done <- true
				return
			}
		}
	}()
	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("send", i)
	}
	fmt.Println("run there")
	close(jobs)
	<-done
}

func RangeChannel() {
	c1 := make(chan int, 2)
	for i := 0; i < 2; i++ {
		c1 <- i
	}
	//close(c1)
	for i := range c1 {
		fmt.Println(i)
	}
}

func TickerChannel() {
	ticker1 := time.NewTicker(time.Second)
	go func() {
		for i := range ticker1.C {
			fmt.Println(i)
		}
	}()

	time.Sleep(time.Second * 3)
	ticker1.Stop()
	fmt.Println("stop")
}

func WorkerChannel() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	for i := 0; i < 3; i++ {
		go func(id int, jobs <-chan int, results chan<- int) {
			for job := range jobs {
				fmt.Println("worker", id, "process", job)
				time.Sleep(time.Second)
				results <- job
			}
		}(i, jobs, results)
	}

	for i := 0; i < 9; i++ {
		jobs <- i
	}
	close(jobs)
	for a := 1; a <= 9; a++ {
		<-results
	}
}

func Atomic() {
	// 我们将使用一个无符号整形数来表示（永远是正整数）这个计数器。
	var ops uint64 = 0

	// 为了模拟并发更新，我们启动 50 个 Go 协程，对计数
	// 器每隔 1ms （译者注：应为非准确时间）进行一次加一操作。
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 使用 `AddUint64` 来让计数器自动增加，使用
				// `&` 语法来给出 `ops` 的内存地址。
				atomic.AddUint64(&ops, 1)

				// 允许其它 Go 协程的执行
				runtime.Gosched()
			}
		}()
	}

	// 等待一秒，让 ops 的自加操作执行一会。
	time.Sleep(time.Second)

	// 为了在计数器还在被其它 Go 协程更新时，安全的使用它，
	// 我们通过 `LoadUint64` 将当前值得拷贝提取到 `opsFinal`
	// 中。和上面一样，我们需要给这个函数所取值的内存地址 `&ops`
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)

}
