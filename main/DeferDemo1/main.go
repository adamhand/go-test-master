package main

import (
	"fmt"
)

func finished()  {
	fmt.Println("function has finished")
}

func findLagest(nums []int)  {
	defer finished()
	max := nums[0]
	for _, v := range nums{
		if v > max{
			max = v
		}
	}
	fmt.Println("the largest number is ", max)
}

func main()  {
	nums := []int{1,2,3,4,5,6,7,8,9}
	findLagest(nums)
}
