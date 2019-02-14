package main

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author)fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title string
	content string
	author
}

func (p post)details()  {
	fmt.Println("title: ", p.title)
	fmt.Println("content: ", p.content)
	//一旦结构体内嵌套了一个结构体字段，Go 可以使我们访问其嵌套的字段，好像这些字段属于外部结构体一样。
	fmt.Println("author: ", p.author.fullName()) //可替换为p.fullName()
	fmt.Println("bio: ", p.author.bio) //可替换为p.bio
}

func main()  {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post1.details()
}


