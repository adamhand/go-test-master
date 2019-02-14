package main

import "fmt"

type SalaryCalculator interface {
	CalculatorSalary() int
}

type Permanent struct {
	empId int
	basicPay int
	pf int
}

type Contract struct {
	empId int
	basicPay int
}

func (p Permanent) CalculatorSalary() int {
	return p.basicPay + p.pf
}

func (c Contract) CalculatorSalary() int {
	return c.basicPay
}

func totalExpense(s []SalaryCalculator)  {
	expense := 0

	for _, v := range s{
		expense += v.CalculatorSalary()
	}
	fmt.Printf("total expense is : $%d", expense)
}

func main()  {
	p1 := Permanent{1, 3000, 30}
	p2 := Permanent{2, 2000, 20}
	c1 := Contract{3, 4000}

	employees := []SalaryCalculator{p1, p2, c1}
	totalExpense(employees)
}
