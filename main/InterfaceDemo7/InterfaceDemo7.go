package main

import "fmt"

type Describer interface {
	describe()
}

type Person struct {
	name string
	age int
}

func (p Person)describe()  {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{})  {
	switch v := i.(type) {
	case Describer:
		v.describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func main(){
	findType("fdsfds")
	p := Person{
		name : "zhangsan",
		age : 30,
	}
	findType(p)
}
