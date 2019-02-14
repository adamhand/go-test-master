package main

import (
	"fmt"
	"sync"
)

var i = 0

func increment(wg *sync.WaitGroup, lock *sync.Mutex)  {
	lock.Lock()
	i = i + 1
	lock.Unlock()
	wg.Done()
}

func main()  {
	var wg sync.WaitGroup
	var lock sync.Mutex

	for x := 0; x < 1000; x++{
		wg.Add(1)
		go increment(&wg, &lock)
	}
	wg.Wait()
	fmt.Printf("result is %d", i)
}
