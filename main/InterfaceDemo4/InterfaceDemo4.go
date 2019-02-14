package main

import "fmt"

func describe(i interface{}){
	fmt.Printf("type = %T, value = %v\n", i, i)
}

func main()  {
	s := "hello world"
	describe(s)

	i := 10
	describe(i)

	st := struct {
		name string
	}{
		name : "zhangsan",
	}
	describe(st)
}
