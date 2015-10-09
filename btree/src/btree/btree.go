package main

import (
	//"fmt"
	"methods"
)

func main() {
	root := NewNode(nil, nil)
	var it Ititer
	it = root
	it.SetData("root node")
	a := NewNode(nil, nil)
	a.SetData("left node")
	al := NewNode(nil, nil)
	al.SetData(100)
	ar := NewNode(nil, nil)
	ar.SetData(3.14)
	a.Left = al
	a.Right = ar
	b := NewNode(nil, nil)
	b.SetData("right node")
	root.Left = a
	root.Right = b
	root.PrintBT()
}
