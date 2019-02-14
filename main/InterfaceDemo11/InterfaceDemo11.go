package main

import "fmt"

type Describer interface {
	describe()
}

func main()  {
	var d1 Describer
	if d1 == nil{
		fmt.Printf("d1 is nil and has type %T value %v", d1, d1)
	}
}
