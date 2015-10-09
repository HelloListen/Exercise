package link

import (
	"fmt"
)

type Student struct {
	Id   int
	Name string
}

type Node struct {
	Student
	Next *Node
}

func (head *Node) Create() *Node {
	head = nil
	return head
}

func (p *Node) PrintLink() {
	for p != nil {
		fmt.Printf("%d,%s\n", p.Id, p.Name)
		p = p.Next
	}
}

func (newNode *Node) Insert(head *Node) *Node {
	var p0, p1, p2 *Node
	p0 = newNode
	p1 = head
	if head == nil {
		head = p0
		p0.Next = nil
	} else {
		for (p0.Id > p1.Id) && p1.Next != nil {
			p2 = p1
			p1 = p1.Next
		}
		if p0.Id <= p1.Id {
			if head == p1 {
				head = p0
			} else {
				p2.Next = p0
				p0.Next = p1
			}
		} else {
			p1.Next = p0
			p0.Next = nil
		}
	}
	return head
}

func (delNode *Node) Delete(head *Node) *Node {
	var p1, p2 *Node
	if head == nil {
		fmt.Println("List nil!")
		goto End
	}
	p1 = head
	for delNode.Id != p1.Student.Id && p1.Next != nil {
		p2 = p1
		p1 = p1.Next
	}
	if delNode.Id == p1.Student.Id {
		if p1 == head {
			head = p1.Next
		} else {
			p2.Next = p1.Next
		}
		fmt.Printf("Delete %d\n", delNode.Id)
	} else {
		fmt.Printf("Node %d not been found!\n", delNode.Id)
	}
End:
	return head
}
