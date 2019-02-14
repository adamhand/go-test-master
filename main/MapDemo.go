package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	personSalary := map[string]int{
		"bob":1000,
		"alice":1200,
	}

	personSalary["mike"] = 1300
	newEmp := "joe"
	value, ok := personSalary[newEmp]
	if ok == true{
		fmt.Println("salary of ", newEmp, "is ", value)
	}else {
		fmt.Println(newEmp," is not found")
	}

	for k, v := range personSalary{
		fmt.Printf("personSalary[%s] = %d\n", k, v)
	}

	delete(personSalary, "bob")
	for k, v := range personSalary{
		fmt.Printf("personSalary[%s] = %d\n", k, v)
	}

	s := "sfsdfasfas"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}
