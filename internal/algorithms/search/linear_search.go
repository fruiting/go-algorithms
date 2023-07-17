package search

type LinearSearch struct {
	complexity string
}

func NewLinearSearch() *LinearSearch {
	return &LinearSearch{
		complexity: "O(n)",
	}
}

func (s *LinearSearch) Complexity() string {
	return s.complexity
}

func (s *LinearSearch) Search(arr []int, val int) int {
	for k, v := range arr {
		if v == val {
			return k
		}
	}

	return -1
}
