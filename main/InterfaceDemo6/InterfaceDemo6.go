package main

import "fmt"

func findType(i interface{})  {
	switch i.(type) {
	case int:
		fmt.Printf("i am an int and my value is %d,\n", i.(int))
	case string:
		fmt.Printf("i am a string and my value is %s\n", i.(string))
	default:
		fmt.Printf("unknown type\n")
	}
}

func main()  {
	findType(33)
	findType("wewe")
	findType(33.44)
}
