package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
工作和结果两个结构体
 */
type Job struct {
	id int
	randomno int
}

type Results struct {
	job Job
	sumofdigits int
}

/*
存放作业和结果的信道
 */
var jobs = make(chan Job, 10)
var results = make(chan Results, 10)

/*
用来计算结果的函数
 */
func digits(number int) int {
	sum := 0
	no := number
	for no != 0{
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(100 * time.Millisecond)
	return sum
}

/*
用来计算的协程，将结果写入Results信道中
 */
func worker(wg * sync.WaitGroup)  {
	for job := range jobs{
		output := Results{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

/*
创建工作池
 */
func createWorkerPool(noOfWorkers int)  {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++{
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

/*
创建工作
 */
func alloate(noOfJobs int)  {
	for i := 0; i < noOfJobs; i++{
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

/*
打印结果
 */
func result(done chan bool)  {
	for r := range results{
		fmt.Printf("job id %d, input num %d, result %d\n", r.job.id, r.job.randomno, r.sumofdigits)
	}
	done <- true
}

func main()  {
	startTime := time.Now()
	noOfJobs := 100
	noOfWorkers := 10
	done := make(chan bool)

	go alloate(noOfJobs)
	go result(done)
	go createWorkerPool(noOfWorkers)
	<- done

	endTime := time.Now()
	diffTime :=endTime.Sub(startTime)
	fmt.Println("time costs is ", diffTime.Seconds(), " s")
}