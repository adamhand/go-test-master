package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width int
}

func (r rect)area(wg *sync.WaitGroup)  {
	defer wg.Done()
	if r.length < 0{
		fmt.Printf("length should greater than zero\n")
		return
	}

	if r.width < 0{
		fmt.Printf("width should greater than zero\n")
		return
	}

	area := r.length * r.width
	fmt.Printf("area is %d\n", area)
}

func main()  {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}