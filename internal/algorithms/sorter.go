package algorithms

type Sorter interface {
	Complexity() string
	Sort(arr []int) []int
}
