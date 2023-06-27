package datastructs

type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int) *Node {
	return &Node{
		Val: val,
	}
}

type List struct {
	head *Node
}

func NewList() *List {
	return &List{}
}

func (l *List) Append(val int) {
}

func (l *List) Remove(val int) {
}
