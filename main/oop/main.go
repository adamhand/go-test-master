package main

import "firstgo/main/oop/employee"

func main()  {
	e := employee.Employee{
		FirstName : "Bob",
		LastName : "Adam",
		TotalLeaves: 100,
		LeavesToken: 10,
	}
	e.LeavesRemaining()
}
