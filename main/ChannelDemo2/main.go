package main

import "fmt"

func hello(done chan bool)  {
	fmt.Printf("hello go channel\n")
	done <- true
}

func main()  {
	done := make(chan bool)
	go hello(done)
	<- done
	fmt.Printf("main function")
}
