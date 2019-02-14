package main

import (
	"fmt"
	"runtime"
)

func printValue(v int)  {
	fmt.Printf("value of v is %d\n", v)
}

func main()  {
	v := 5
	defer printValue(v)
	v = 10
	fmt.Printf("value of v is %d\n", v)
}
