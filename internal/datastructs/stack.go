package datastructs

import "fmt"

type Stack struct {
	values []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Run() {
	fmt.Println("Filling stack")

	for i := 0; i < 10; i++ {
		s.values = append(s.values, i)
		fmt.Println(s.values)
	}

	fmt.Println("Filled stack")
	fmt.Println(s.values)
	fmt.Println("Clearing stack")

	for i := 10; i >= 0; i-- {
		s.values = s.values[0:i]
		fmt.Println(s.values)
	}
}
