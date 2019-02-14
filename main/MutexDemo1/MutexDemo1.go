package main

import (
	"fmt"
	"sync"
)

var i = 0

func incresment(wg *sync.WaitGroup)  {
	i = i + 1
	wg.Done()
}

func main()  {
	var wg sync.WaitGroup

	for x := 0; x < 1000; x++{
		wg.Add(1)
		go incresment(&wg)
	}
	wg.Wait()
	fmt.Printf("result is %d", i)
}
