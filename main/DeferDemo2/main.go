package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func (p person)fullName()  {
	fmt.Printf("%s", p.firstName + p.lastName)
}

func main()  {
	p := person{
		firstName:"Bob",
		lastName:"Steven",
	}
	defer p.fullName()
	fmt.Print("welcome ")
}
