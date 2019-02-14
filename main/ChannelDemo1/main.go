package main

import "fmt"

func main()  {
	var a chan int
	if a == nil{
		fmt.Printf("channel is nil, going to define it\n")
		a = make(chan int)
		fmt.Printf("type of a is %T", a)
	}
}
