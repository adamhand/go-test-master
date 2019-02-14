package main

import "fmt"

func calSquares(number int, sqareop chan int) {
	sum := 0
	for number != 0{
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	sqareop <- sum
}

func calCubes(number int, cubeop chan int)  {
	sum := 0
	for number != 0{
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main()  {
	number := 123
	squarech := make(chan int)
	cubech := make(chan int)
	go calCubes(number, cubech)
	go calSquares(number, squarech)
	squares, cubes :=  <- squarech, <- cubech
	fmt.Printf("final results: %d, %d", squares, cubes)
}
