package main

import "fmt"

//func main(){
//	find(89, 89, 90, 95)
//	find(45, 56, 67, 45, 90, 109)
//	find(78, 38, 56, 98)
//	find(87)
//}
//
//func find(num int, nums ...int){
//	fmt.Printf("type of nums is %T\n", nums)
//	isFind := false
//
//	for i, v := range nums{
//		if v == num{
//			fmt.Println(num, "found at index ", i, "in", nums)
//			isFind = true
//		}
//	}
//	if !isFind{
//		fmt.Println(num, "not found in", nums)
//	}
//	fmt.Printf("\n")
//}

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func main() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}
