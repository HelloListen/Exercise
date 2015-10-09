package btree

import (
	"fmt"
)

type Node struct {
	Left  *Node
	Data  interface{}
	Right *Node
}

type Initer interface {
	SetData(data interface{})
}

type Operater interface {
	PrintBT()
	Depth() int
	LeafCount() int
}

type Order interface {
	PreOrder()
	InOrder()
	PostOrder()
}

func (n *Node) SetData(data interface{}) {
	n.Data = data
}

func (n *Node) PrintBT() {
	PrintBT(n)
}

func (n *Node) Depth() int {
	return Depth(n)
}

func (n *Node) LeafCount() int {
	return LeafCount(n)
}

func (n *Node) PreOrder() {
	PreOrder(n)
}

func (n *Node) InOrder() {
	InOrder(n)
}

func (n *Node) PostOrder() {
	PostOrder(n)
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func PrintBT(n *Node) {
	if n != nil {
		fmt.Printf("%v", n.Data)
		if n.Left != nil || n.Right != nil {
			fmt.Printf("(")
			PrintBT(n.Left)
			if n.Right != nil {
				fmt.Printf(",")
			}
			PrintBT(n.Right)
			fmt.Printf(")")
		}
	}
}

func Depth(n *Node) int {
	var depleft, depright int
	if n == nil {
		return 0
	} else {
		depleft = Depth(n.Left)
		depright = Depth(n.Right)
		if depleft > depright {
			return depleft + 1
		} else {
			return depright + 1
		}
	}
}

func LeafCount(n *Node) int {
	if n == nil {
		return 0
	} else if (n.Left == nil) && (n.Right == nil) {
		return 1
	} else {
		return (LeafCount(n.Left) + LeafCount(n.Right))
	}
}

func PreOrder(n *Node) {
	if n != nil {
		fmt.Printf("%v", n.Data)
		PreOrder(n.Left)
		PreOrder(n.Right)
	}
}

func InOrder(n *Node) {
	if n != nil {
		PreOrder(n.Left)
		fmt.Printf("%v", n.Data)
		PreOrder(n.Right)
	}
}

func PostOrder(n *Node) {
	PreOrder(n.Left)
	PreOrder(n.Right)
	fmt.Printf("%v", n.Data)
}
