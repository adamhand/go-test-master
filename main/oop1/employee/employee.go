package employee

import "fmt"

type employee struct {
	firtname string
	lastname string
	totalleaves int
	leavestoken int
}

func New(firstName string, lastName string, totalLeaves int, leavesToken int) employee  {
	e := employee{firstName,lastName, totalLeaves, leavesToken}
	return e
}

func (e employee)LeavesRamain()  {
	fmt.Printf("%s %s has %d leaves remain", e.firtname, e.lastname, (e.totalleaves - e.leavestoken))
}