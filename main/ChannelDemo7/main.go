package main

import (
	"fmt"
	"time"
)

func write(channel chan int)  {
	for i := 0; i < 5; i++{
		channel <- i
		fmt.Println("successfully write ", i, " to channel")
	}
	close(channel)
}

func main()  {
	ch := make(chan int, 2)
	go write(ch)

	time.Sleep(2 * time.Second)
	for v := range ch{
		fmt.Println("read value ",v, " from channel")
		time.Sleep(2 * time.Second)
	}
}
