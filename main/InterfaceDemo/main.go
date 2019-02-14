package main

import "fmt"

type Invoker interface {
	call(interface{})
}

type Structure struct {

}

type FuncCall func(interface{})

func (s *Structure)call(p interface{})  {
	fmt.Println("from structure", p)
}

func (f FuncCall)call(p interface{})  {
	f(p)
}

func main() {
	var invoker Invoker

	s := new(Structure)

	invoker = s

	invoker.call("hello")

	invoker = FuncCall(func(v interface{}) {
		fmt.Println("from function", v)
	})

	invoker.call("hello")
}
