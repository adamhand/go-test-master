package main

import "fmt"

func main()  {
	name := "stevem bob"
	fmt.Printf("name is: %s\n", name)
	fmt.Printf("reverse name is: ")

	for _, v := range []rune(name){
		defer fmt.Printf("%c", v)
	}
}
