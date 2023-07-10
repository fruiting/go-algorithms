package datastructs

import "fmt"

type Queue struct {
	values []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (s *Queue) Run() {
	fmt.Println("Filling queue")

	for i := 0; i < 10; i++ {
		s.values = append(s.values, i)
		fmt.Println(s.values)
	}

	fmt.Println("Filled queue")
	fmt.Println(s.values)
	fmt.Println("Clearing queue")

	for len(s.values) > 0 {
		s.values = s.values[1:]
		fmt.Println(s.values)
	}
}
