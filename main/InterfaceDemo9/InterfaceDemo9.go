package main

import "fmt"

type SalaryCalculator interface {
	displaySalary()
}

type LeaveCalculator interface {
	calculatorLeavesLeft() int
}

type Employee struct {
	firstName string
	lastName string
	basicPay int
	pf int
	totalLeaves int
	leavesToken int
}

func (e Employee)displaySalary()  {
	fmt.Printf("%s %s has salary $%d\n", e.firstName, e.lastName, (e.pf + e.basicPay))
}

func (e Employee)calculatorLeavesLeft() int {
	return e.totalLeaves - e.leavesToken
}

func main()  {
	e := Employee{
		firstName:"bob",
		lastName:"steven",
		basicPay:3000,
		pf:2000,
		totalLeaves:10,
		leavesToken:3,
	}
	var s SalaryCalculator = e
	s.displaySalary()
	var l LeaveCalculator = e
	fmt.Printf("\nleaves left = %d", l.calculatorLeavesLeft())
}
