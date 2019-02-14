package main

import (
	"fmt"
	"time"
)

func server1(ch chan string)  {
	ch <- "server1"
}

func server2(ch chan string)  {
	ch <- "server 2"
}

func main()  {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)
	time.Sleep(100 * time.Millisecond)

	select {
	case out1 := <- ch1:
		fmt.Println("received value ", out1)
	case out2 := <- ch2:
		fmt.Println("received value ", out2)
	}
}
