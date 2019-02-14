package main

import "fmt"

type SalaryCalculator interface {
	displySalary()
}

type LeavesCalculator interface {
	calculatorLeavesLeft() int
}

type EmployeeOprations interface {
	SalaryCalculator
	LeavesCalculator
}

type Employee struct {
	firstName string
	lastName string
	basicPay int
	pf int
	totalLeaves int
	leavesToken int
}

func (e Employee)displySalary()  {
	fmt.Printf("%s %s has %d salary\n", e.firstName, e.lastName, (e.pf + e.basicPay))
}

func (e Employee)calculatorLeavesLeft() int {
	return e.totalLeaves - e.leavesToken
}

func main()  {
	e := Employee{
		firstName:"bob",
		lastName:"steven",
		basicPay:4000,
		pf:1000,
		totalLeaves:10,
		leavesToken:5,
	}
	var empOp EmployeeOprations = e
	empOp.displySalary()
	fmt.Printf("\nleaves left = %d", empOp.calculatorLeavesLeft())
}
