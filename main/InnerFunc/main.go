package main

import "fmt"

//匿名函数
func main()  {
	res := intSeq()

	fmt.Print(res())
	fmt.Print(res())
	fmt.Print(res())
	fmt.Print(res())

	res1 := intSeq()
	fmt.Print(res1())
}

func intSeq() (func() (int))  {
	i := 0

	return func() int {
		i += 1
		return i
	}
}