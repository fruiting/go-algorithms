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
	newNode := &Node{Val: val}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode
}

func (l *List) Remove(val int) {
	if l.head == nil {
		return
	}

	if l.head.Val == val {
		l.head = l.head.Next
		return
	}

	curr := l.head
	for curr.Next != nil && curr.Next.Val != val {
		curr = curr.Next
	}

	if curr.Next != nil {
		curr.Next = curr.Next.Next
	}
}
