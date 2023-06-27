package algorithms

type Searcher interface {
	Complexity() string
	Search(sortedArr []int, val int) int
}
