package main

import (
	"link"
)

func main() {
	var head *link.Node
	stu1 := link.Node{link.Student{100, "LiMing"}, nil}
	stu2 := link.Node{link.Student{101, "ZhangXiao"}, nil}
	stu3 := link.Node{link.Student{102, "Listen"}, nil}
	stu4 := link.Node{link.Student{103, "Mike"}, nil}
	head = head.Create()
	head = stu1.Insert(head)
	head = stu2.Insert(head)
	head = stu3.Insert(head)
	head = stu4.Insert(head)
	head.PrintLink()
	head = stu3.Delete(head)
	head.PrintLink()
}
