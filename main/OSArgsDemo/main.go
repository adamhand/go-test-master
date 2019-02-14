package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main()  {
	//if len(os.Args) != 0 {
	//	fmt.Println(os.Args[0]) // args 第一个片 是文件路径
	//}
	//fmt.Println(os.Args[1])  //用户输入的参数

	//test1()
	//test2()
	test3()
}

/**
统计出现次数大于1的值，使用ctrl + d结束程序
输入：
1
2
3
4
5
1
1
1
^D
输出：
4	1
 */
func test1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan(){
		if input.Text() == "end"{
			//break
		}
		counts[input.Text()]++
	}

	for line, n := range counts{
		if n > 1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func test2()  {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	}else {
		for _, arg := range files{
			f, err := os.Open(arg)
			if err != nil{
				fmt.Fprint(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts{
		if n > 1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()]++
	}
}

/**
需要传入一个文件作为ReadFile的参数
 */
func test3()  {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:]{
		data, err := ioutil.ReadFile(filename)
		if err != nil{
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n"){
			counts[line]++
		}
	}
	for line, n := range counts{
		if n > 1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}