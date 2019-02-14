package main

import "fmt"

type Describer interface {
	describe()
}

type Person struct {
	name string
	age int
}

type Address struct {
	state string
	country string
}

func (p Person)describe()  {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

func (a *Address)describe()  {
	fmt.Printf("state %s country %s\n", a.state, a.country)
}

func main()  {
	var d1 Describer
	p1 := Person{"bob", 22}
	d1 = p1
	d1.describe()

	p2 := Person{"alice", 33}
	d1 = &p2
	d1.describe()

	var d2 Describer
	a := Address{"shanchong", "China"}
	d2 = &a
	d2.describe()
}
