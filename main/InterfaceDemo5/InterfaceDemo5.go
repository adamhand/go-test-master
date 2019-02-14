package main

import "fmt"

func assert(i interface{})  {
	//s := i.(int)
	//fmt.Println(s)

	v, ok := i.(int)
	//fmt.Println(v, ok)
	if ok{
		fmt.Println(v)
	}else{
		fmt.Println("type is not correct, the correct type is %T ", v)
	}
}

func main()  {
	var i interface{} = 55
	assert(i)

	var s interface{} = "hahahah"
	assert(s)
}
