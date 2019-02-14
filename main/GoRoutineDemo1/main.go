package main

import (
	"fmt"
	"time"
)

func hello()  {
	fmt.Printf("hello world go routine\n")
}

func main()  {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Printf("main function\n")
}
