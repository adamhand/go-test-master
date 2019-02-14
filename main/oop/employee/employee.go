package employee

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	TotalLeaves int
	LeavesToken int
}

func (e Employee)LeavesRemaining()  {
	fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesToken))
}
