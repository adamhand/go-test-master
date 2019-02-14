package main

import (
	"fmt"
	"strings"
)

func main()  {
	var num int = 12
	fmt.Print(num, "\n")

	var num1, num2, num3 int  = 1, 2, 3

	fmt.Print(num1, num2, num3, "\n")

	num4, num5, num6 := 4, 5, 6
	fmt.Print(num4, num5, num6, "\n")

	var result int = add(num5, num6)

	fmt.Print("hello world", result, "\n")

	var string1 string = "hahhahahh"

	fmt.Printf(string1+"\n")
	fmt.Printf("%s", string1)

	for i := 0; i < len(string1); i++ {
		fmt.Printf("%x ", string1[i])
	}

	stringarr := []string{"nihaoa", "niyehao"}

	string2 := strings.Join(stringarr, " ")

	fmt.Printf(string2)

	numdemo()

	pointernum()

	doublepointer()

	functionpointer()

	structdemo()

	ifelsedemo(num, num1)

	sliecedemo()

	rangedemo()

	var i int=  fibonaci(5)
	fmt.Print(i)
}

func add(a int, b int) (int)  {
	var c int

	c = a + b

	return c
}

func numdemo() {
	var nums [10]int

	var nums1 = [5]int{90, 91, 92, 93, 94}

	for i := 0; i < len(nums); i++ {
		nums[i] = i
	}

	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}

	for i := 0; i < len(nums1); i++ {
		fmt.Print(nums1[i], " ")
	}
}

func pointernum() {
	var num  =	[5]int{1, 2, 3, 4, 5}

	var p [5]*int

	for i := 0; i < len(num); i++ {
		p[i] = &num[i]
	}

	for i := 0; i < len(num); i++ {
		fmt.Print(*p[i], " ")
	}
}

func doublepointer() {
	var num int = 100
	var p *int
	var pp **int

	p = &num
	pp = &p

	fmt.Print(num)
	fmt.Print(*p)
	fmt.Print(**pp)
}

func functionpointer() {
	var a = 100
	var b = 200
	
	fmt.Print(a, " ", b)

	swap(&a, &b)

	fmt.Print(a, " ", b)
}

func swap(a *int, b *int) {
	var temp int

	temp = *a
	*a = *b
	*b = temp
}

type book struct {
	name string
	author string
	id string
	price int
}

func structdemo() {
	fmt.Print("\n")

	var book1 book
	var book2 book

	book1 = book{"Java并发编程实战", "noone", "1", 100}

	book2.name = "深入理解Java虚拟机"
	book2.author = "周志华"
	book2.id = "2"
	book2.price = 90

	fmt.Print(book1.name, " ", book1.author, " ", book1.id, " ", book1.price, "\n")
	fmt.Print(book2.name, " ", book2.author, " ", book2.id, " ", book2.price, "\n")
}

func ifelsedemo(a int, b int) {
	if a < b {
		fmt.Print(a)
	}else {
		fmt.Print(b)
	}
}

func sliecedemo()  {
	var slicenum = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	subslice1 := slicenum[2:3]

	fmt.Print(subslice1)

	slicenum = append(slicenum, 10, 11, 12)

	fmt.Print(slicenum, len(slicenum), cap(slicenum))

	slicenum1 := make([]int, len(slicenum), cap(slicenum)*2)
	copy(slicenum1, slicenum)

	fmt.Print(slicenum1, len(slicenum1), cap(slicenum1))
}

func rangedemo()  {
	var slicenum = []int{0, 1, 2, 3, 4, 5}

	for i := range slicenum{
		fmt.Print(slicenum[i], " ")
	}

	demoMap := map[string] string {"France":"Paris", "Italy":"Rome", "Chana":"Beijing"}

	for i := range demoMap{
		fmt.Print("capital of ", i, " is ", demoMap[i], "\n")
	}

	for i, j := range demoMap{
		fmt.Print("capital of ", i, " is ", j, "\n")
	}
}

func fibonaci(i int) (int) {
	if i == 0{
		return 0
	}
	if i == 1{
		return 1
	}

	return fibonaci(i - 1) + fibonaci(i - 2)
}