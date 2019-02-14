package main

import (
	"fmt"
	"time"
)

func alphabets()  {
	for i := 'a'; i <= 'e'; i++{
		time.Sleep(25 * time.Millisecond)
		fmt.Printf("%c ",i)
	}
}

func numbers()  {
	for i := 0; i <= 5; i++{
		time.Sleep(40 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main()  {
	go alphabets()
	go numbers()
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("main function")
}
