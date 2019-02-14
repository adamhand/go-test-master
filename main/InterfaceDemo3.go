package main

import "fmt"

type Test interface {
	Tester()
}

type MyFloat float64

func (m MyFloat)Tester()  {
	fmt.Println(m)
}

func describe(t Test){
	//t既包含一个类型也包含一个值
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func main(){
	var t Test
	f := MyFloat(43.90)
	t = f
	describe(t)
	t.Tester()
}
