package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, group *sync.WaitGroup)  {
	fmt.Println("start go routine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("go routine %d end\n", i)
	group.Done()
}

func main()  {
	no := 3
	var group sync.WaitGroup
	for i := 0; i < no; i++{
		group.Add(1)
		go process(i, &group)
	}
	group.Wait()
	fmt.Println("all go routines finished")
}
